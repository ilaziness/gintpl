package main

import (
	"github.com/ilaziness/gintpl/internal/app/web"
	"github.com/ilaziness/gokit/storage/mysql"
	"github.com/spf13/cobra"
	"gorm.io/gen"
)

// gorm相关的工具

var gormCmd = &cobra.Command{
	Use:   "gorm",
	Short: "gorm generator tool",
	Long:  `execute gorm generator tool`,
}

var gormDAOCmd = &cobra.Command{
	Use:   "dao",
	Short: "generator dao query and model code from database",
	Long: `generator dao query and model code from database. For example:

go run ./cmd gorm dao -o ./internal/query

Inside the directory internal, a query directory and a model directory will be generated,
which will respectively store the generated query code and model code `,
	Run: func(cmd *cobra.Command, _ []string) {
		op, _ := cmd.Flags().GetString("out_path")
		mysql.GenerateDAO(web.Config.Db, op, func(g *gen.Generator) {
			// 自定义对应模型
			g.ApplyBasic(g.GenerateAllTable()...)
		})
	},
}

func init() {
	gormCmd.AddCommand(gormDAOCmd)
	gormDAOCmd.Flags().StringP("out_path", "o", "", "the directory name to save query code file, default is ./query")
}
