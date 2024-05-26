Feature: reporting data service

    Scenario: all existing patients are shwon in patient list view
        Given head into the reporting data home page
        When any patients exist in the database
        Then all patients will be seen on the browser
