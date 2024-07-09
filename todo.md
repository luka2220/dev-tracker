# TODOs
- [1] Add CLI commands [✅]:
    * Name the base CLI command [devtasks]?
    - Create a new development board => [devtasks init]
    - Intract with the tasks and development boards => [devtasks]

- [2] Create the init TUI program [🕐]:
    * Creates the project initalization cli for creating new development boards and storing them in a database
    - Prompt the user to create a new development board (completed) ✅
    - Validate user input (completed) ✅
    - Warn the user of input error (completed) ✅
    - Store in a sqlite database (completed) ✅
        * Create the connection to the sqlite database - done
        * Create the table and schema for what each board can hold - done
        * Create an entry in the db for the newly created board and it's options - done
    - Store users initialization board into the databse (completed) ✅
        * Take the inputs and format them correctly to be stored in the db - done
        * Update the project model state to make it clear when we should be storing to the db - done
        * Respond to the user with any errors that occur with proper info - done
        * Test to validate the board was correctly stored in the database - done
        * Create a log entry to db.log whenever a new board record is created in the db - done
        * Create a sepreate db-error.log to write database errors to - done

- [3] Add menu for display the options for operations below [🕐]:
    * Add a menu to display all of the options below and a text input for the user to choose an option
    - Set the currently active board from the database to operate on (completed) ✅
    - The menu options should be:
        * [1] create new task -> create a new task for the current development board
        * [2] update/check task -> update or check a task for the current development board
        * [3] change development board -> switch to a different board
        * [4] show the tasks detail -> display the tasks detail
        * [5] delete development board -> delete the selected board
