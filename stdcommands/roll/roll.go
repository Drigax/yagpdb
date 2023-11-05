package roll

import (
	"strings"

	"github.com/botlabs-gg/yagpdb/v2/commands"
	"github.com/botlabs-gg/yagpdb/v2/lib/dcmd"
	"github.com/botlabs-gg/yagpdb/v2/lib/dice"
)

var Command = &commands.YAGCommand{
	CmdCategory:     commands.CategoryFun,
	Name:            "Roll",
	Description:     "Roll dice, using rpg dice syntax. Rolls 1d20 by default.",
	LongDescription: "Example: `-roll 2d6`",
	Arguments: []*dcmd.ArgDef{
		{Name: "Dice_Notation", Type: dcmd.String},
	},
	DefaultEnabled:      true,
	SlashCommandEnabled: true,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		output := ""
		input := "1d20"

		if data.Args[0].Value != nil {
			input = data.Args[0].Str()
		}

		r, _, err := dice.Roll(strings.ToLower(input))
		output = r.String()
		if err != nil {
			return err.Error(), nil
		}

		if len(output) > 100 {
			output = output[:100] + "..."
		} else {
			output = strings.TrimSuffix(output, "([])")
		}
		return ":game_die: Rolling " + input + ": " + output, nil
	},
}
