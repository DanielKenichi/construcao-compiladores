#!/bin/bash
# Generate a random string
RANDOM_STR=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 10 | head -n 1)

# Append 'temp' to the random string
FOLDER_NAME="${RANDOM_STR}temp"

# Create the full path for the new directory
TEMP_FOLDER="${PWD}/${FOLDER_NAME}"

# Create the new directory
mkdir "${TEMP_FOLDER}"

java -jar $PWD/compiladores-corretor-automatico-1.0-SNAPSHOT-jar-with-dependencies.jar "${PWD}/$1" gcc $TEMP_FOLDER $PWD/casos-de-teste "790837 800345" "$1"
