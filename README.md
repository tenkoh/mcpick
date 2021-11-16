# mcpick
Light weight Terminal User Interface (TUI) to pick material colors.
You do NOT need to take your hands off the keyboard to pick colors.

You can choose output style, web format (#ffffff) or rgb format (255, 255, 255).

## Getting started
1. Save the binary file `mcpick` in any directory.
1. Add the path to PATH.

## Usage
This app just outputs color code to the terminal. You can combinate this app with others like `pbcopy` in MacOS or `clip` in Windows.

### Example: MacOS
The example below copies color code to the clipboard. The most convenient way is make alias in your `.bashrc` or `.zshrc`.

```shell
mcpick | pbcopy
```

### Example: Windows
The example below copies color code to the clipboard. The most convenient way is make `.bat` file and add the path to `PATH`.

```shell
mcpick | clip
```