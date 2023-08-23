package cmd

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/borkborkbork/pkg/clparser"
)

// renderCmd represents the render command
var renderCmd *cobra.Command

func NewRenderCommand() *cobra.Command {
	renderCmd = &cobra.Command{
		Use:   "render",
		Short: "A brief description of your command",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("arg count: ", len(args))
			return render("example.cook")
		},
	}

	return renderCmd
}

const OFFSET_INDENT = 4

func render(fname string) error {
	recipe, err := clparser.ParseFile(fname)
	if err != nil {
		return fmt.Errorf("parsing recipe from file: %w", err)
	}
	printRecipe(*recipe, os.Stdout)

	return nil
}

func collectIngredients(steps []clparser.Step) []clparser.Ingredient {
	var result []clparser.Ingredient
	for i := range steps {
		result = append(result, steps[i].Ingredients...)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result
}

func coollectCookware(steps []clparser.Step) []string {
	var result []string

	for i := range steps {
		for j := range steps[i].Cookware {
			result = append(result, steps[i].Cookware[j].Name)
		}
	}

	sort.Strings(result)

	return result
}

func formatFloat(num float64, precision int) string {
	fs := fmt.Sprintf("%%.%df", precision)
	s := fmt.Sprintf(fs, num)

	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

func getIngredients(ing []clparser.Ingredient) []string {
	var result []string

	for i := range ing {
		result = append(result,
			fmt.Sprintf("%s: %s %s",
				ing[i].Name,
				formatFloat(ing[i].Amount.Quantity, 2),
				ing[i].Amount.Unit))
	}

	sort.Strings(result)

	return result
}

func printRecipe(recipe clparser.Recipe, out io.Writer) {
	offset := strings.Repeat(" ", OFFSET_INDENT)

	if len(recipe.Metadata) > 0 {
		fmt.Fprintln(out, "Metadata:")

		for k, v := range recipe.Metadata {
			fmt.Fprintf(out, "%s%s: %s\n", offset, k, v)
		}

		fmt.Fprintln(out, "")
	}

	allIngredients := collectIngredients(recipe.Steps)
	if len(allIngredients) > 0 {
		fmt.Fprintln(out, "Ingredients:")

		for i := range allIngredients {
			fmt.Fprintf(out,
				"%s%-30s%s %s\n",
				offset,
				allIngredients[i].Name,
				formatFloat(allIngredients[i].Amount.Quantity, 2),
				allIngredients[i].Amount.Unit)
		}

		fmt.Fprintln(out, "")
	}

	allCookware := coollectCookware(recipe.Steps)
	if len(allCookware) > 0 {
		fmt.Fprintln(out, "Cookware:")

		for i := range allCookware {
			fmt.Fprintf(out, "%s%s\n", offset, allCookware[i])
		}

		fmt.Fprintln(out, "")
	}

	if len(recipe.Steps) > 0 {
		fmt.Fprintln(out, "Steps:")

		for i := range recipe.Steps {
			fmt.Fprintf(out, "%s%2d. %s\n", offset, i+1, recipe.Steps[i].Directions)
			ingredients := "–"
			ing := getIngredients(recipe.Steps[i].Ingredients)
			if len(ing) > 0 {
				ingredients = strings.Join(ing, "; ")
			}

			fmt.Fprintf(out, "%s    [%s]\n", offset, ingredients)
		}
	}
}
