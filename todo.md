# TODOs
- [1] Add CLI commands [✅]:
    * Name the base CLI command [devtasks]?
    - Create a new development board => [devtasks init]
    - Intract with the tasks and development boards => [devtasks]

- [2] Create the init TUI program [🕐]:
    * Create the project initalization cli
    - Prompt the user to create a new development board (completed)
    - Validate user input (progress)
    - Store in a sqlite database

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
