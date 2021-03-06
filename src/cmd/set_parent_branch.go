package cmd

import (
	"github.com/Originate/git-town/src/git"
	"github.com/spf13/cobra"
)

var setParentBranchCommand = &cobra.Command{
	Use:   "set-parent-branch <child_branch> <parent_branch>",
	Short: "Updates a branch's parent",
	Long: `Updates a branch's parent

Updates the parent branch of a feature branch in Git Town's configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		setParentBranch(args[0], args[1])
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := validateMaxArgs(args, 2)
		if err != nil {
			return err
		}
		return git.ValidateIsRepository()
	},
}

func setParentBranch(childBranch, parentBranch string) {
	git.EnsureHasBranch(childBranch)
	git.EnsureHasBranch(parentBranch)
	git.SetParentBranch(childBranch, parentBranch)
	git.DeleteAllAncestorBranches()
}

func init() {
	RootCmd.AddCommand(setParentBranchCommand)
}
