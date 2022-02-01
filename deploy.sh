#!/bin/bash

# utility functions hell yeah
panic(){
    >&2 echo "$@"
    exit
}

# -constants-
INIT_SQL="./config/init.sql"
DIRECOTRIES=( uploads uploads/images uploads/videos)

# config directory needs to exist, and init the ".env" file -> we're reading it and basing all configurations on it
CONFIG_FILE="./config/.env"

if [[ -z $CONFIG_FILE ]]; then
    panic "The .env file was not found!"
fi

for i in "${DIRECOTRIES[@]}"
do
    if [ ! -d $i ]
    then
        mkdir "$i"
    fi
done

line_num=0

while read line
do
    ((line_num++))

    # we create an array that contains the key and the value respectively
    IFS="="
    read -a arr <<< "$line"
    
    # if the array doesn't contain exactly two elements, shit's fucked up, so we exit
    if [ ! ${#arr[@]} -eq 2 ]
    then
        panic "Invalid .env syntax on line $line_num!"
    fi

    key=${arr[0]}
    value=${arr[1]}

    case $key in

        "DB_NAME")
            DB_NAME=$value
            ;;

        "DB_USER")
            DB_USER=$value
            ;;

        "DB_PASSWORD")
            DB_PASSWORD=$value
            ;;

        *)
            continue
            ;;

    esac

done < $CONFIG_FILE

# modify the "init.sql" and "delete_all.sql" files to include USE $DB_NAME; if that line isn't there yet
# stole it from here epically: https://superuser.com/questions/246837/how-do-i-add-text-to-the-beginning-of-a-file-in-bash#246841
first_line=$(head -n 1 $INIT_SQL)
if [[ ! $first_line == *"USE"* ]]
then
    echo "USE $DB_NAME;" | cat - $INIT_SQL > temp && mv temp $INIT_SQL
fi

# initialize the database by running the init.sql file
echo "Running init.sql on the database:"
mysql -u $DB_USER -p < $INIT_SQL

if [ ! $? -eq 0 ]
then
    panic "Database initialization failed!"
fi

go run main.go
