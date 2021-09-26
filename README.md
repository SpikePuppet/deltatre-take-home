# Deltatre Take Home Test
## Overview
This project is an implementation of the Deltatre take home test. This project comprises of two services which give the user the ability to search for keywords. You can currently:
1. Search for keywords that exist in the database
2. Update which keywords are available to search
3. Get metrics on how many times existing keywords have been searched.

## Getting Started
This project is designed to be run via docker (so ensure this is installed), specifically docker compose. To run the project, simply run `docker-compose up -d`, and docker will handle generating everything you need to start.
When you build the project, it will both run tests against the provided code and build the artifacts as required. As such you can choose to build containers individually. For example, if you `cd` into the searchtermsbackend folder and use docker build, it will test and build your solution.
Once this is running, you can query the service at localhost:8080 and use the following endpoints
1. GET - /search-term/:search-term - This will search the database for :search-term, and return a message saying either "Term exists!" or "Term does not exist!"
2. PUT - /update-search-terms/:search-term - This will add a term to the database and will start tracking metrics on how many times it's been searched
3. GET /search-term-metrics - This will get the top five most searched terms. Note, if they haven't been searched the field SearchCount will be omitted, meaning it was 0. The results are returned as an array of objects, with a Term property which defined which term it's reporting on, and SearchCount which reports how many times a term has been searched.

## Potential Refactoring Opportunities
If you wish to see what i would start to improve with more time, feel free to reference the PotentialRefactoringOpportunities.md file.