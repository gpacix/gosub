# gosub
Take var=value files and merge them into a bash-style template with as few surprises as possible.

This differs from envsubst in that the environment is ignored; only the specified variable files are consulted.

The last setting for a variable wins, so a typical use would look like:
`$ gosub defaultvars generalvars specificvars < mytemplate > myoutput`

In general, this utility tries to do what setting bash shell variables (not environment variables) and then generating a heredoc would do.

Variable files look like this:
```
# Lines beginning with # are ignored
  # even if there's leading whitespace (spaces or tabs)
# blank/entirely whitespace lines are also ignored

SOME_VAR=fine
MY_VAR=a value including any spaces
MY_QUOTED_VAR="The double-quotes are not part of the value."
MY_SQUOTED_VAR='The single-quotes are not part of the value.'
MY_VAR_KEEPING_QUOTES="'The outer quotes protect the inner.'"
  DOES_LEADING_SPACE_MATTER=no
CAN_YOU_HAVE_SPACES_AROUND_EQUALS = nope; or rather, the variable's name now ends with a space
NO_PARTIAL_LINE_COMMENTS=this value includes # <-- that hash sign and this text
lowercase_var=work
dash-containing=also work
REDEF=old value
REDEF=last-assigned value
```

Template files look line any text file:
```
# This remains part of the resulting file.
# Substitution happens even in comments (how could it tell?): $SOME_VAR
Here's $MY_VAR.
No quotes in this line, since: $MY_QUOTED_VAR
There are single-quotes in this line, since: $MY_VAR_KEEPING_QUOTES
$DOES_LEADING_SPACE_MATTER problems with leading space in variables file.
If the variable wasn't defined, you get: $BLAHBLAH
If you don't want a substitution, put a backslash before the dollar-sign: \$MY_VAR
Oh, yeah: curly-brace substitution works ${SOME_VAR} (but none of the other bash syntax).
Lowercase variables $lowercase_var.
Dash-containing variables $dash-containing.
If you define a variable multiple times, it uses the $REDEF.
```

Currently, undefined variables ARE NOT replaced with the empty string; the reference remains in the output.
