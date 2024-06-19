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
    - Store in a sqlite database (in progress) ⚠️
        * Create the connection to the sqlite database
        * Create the table and schema for what each board can hold
        * Create an entry in the db for the newly created board and it's options

- [3] Add menu for display the options for operations below [🕐]:
    * Add a menu to deisplay all of the options below and a text input for the user to choose an option
    - Develop custom help menu with styling... may look better visually but be more cluttered?
    - User should input numbers to do an operation
    - The menu options should be:
        * [1] create new task -> create a new task for the current development board
        * [2] update/check task -> update or check a task for the current development board
        * [3] change development board -> switch to a different board
        * [4] show the tasks detail -> display the tasks detail
        * [5] delete development board -> delete the selected board
