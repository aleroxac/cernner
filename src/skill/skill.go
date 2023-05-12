package skill

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type Skill struct {
	Id          uuid.UUID `json:"id"`
	Kind        string    `json:"kind"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Scope       string    `json:"scope"`
	Category    string    `json:"category"`
	Purpose     string    `json:"purpose"`
	Level       int       `json:"level"`
	Priority    int       `json:"priority"`
	Tags        []string  `json:"tags"`
}

var name string
var description string
var scope string
var category string
var purpose string
var level int
var priority int
var tags []string
var file string

func createSkill(cmd *cobra.Command, args []string) error {
	if file != "" {
		file_content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
		}

		file_content_string := string(file_content)
		docs := strings.Split(file_content_string, "---")

		var skills []Skill
		for _, doc := range docs {
			// Skip empty documents
			if strings.TrimSpace(doc) == "" {
				continue
			}

			var new_skill Skill
			if err := yaml.Unmarshal([]byte(doc), &new_skill); err != nil {
				fmt.Println(err)
			}

			newSkill := Skill{
				Id:          uuid.New(),
				Kind:        new_skill.Kind,
				Name:        new_skill.Name,
				Description: new_skill.Description,
				Scope:       new_skill.Scope,
				Category:    new_skill.Category,
				Purpose:     new_skill.Purpose,
				Level:       new_skill.Level,
				Priority:    new_skill.Priority,
				Tags:        new_skill.Tags,
			}

			skills = append(skills, newSkill)
		}
		skill_json_formated, err := json.MarshalIndent(skills, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", skill_json_formated)
		return nil
	} else {
		id := uuid.New()
		kind := "skill"
		newSkill := Skill{
			Id:          id,
			Kind:        kind,
			Name:        name,
			Description: description,
			Scope:       scope,
			Category:    category,
			Purpose:     purpose,
			Level:       level,
			Priority:    priority,
			Tags:        tags,
		}

		skill_json_formated, err := json.MarshalIndent(&newSkill, "", "  ")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%s\n", skill_json_formated)
		return nil
	}
}

var SkillCmd = &cobra.Command{
	Use:   "skill",
	Short: "Add a new skill with name and description",
	RunE:  createSkill,
}

func init() {
	SkillCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the skill")
	SkillCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the skill")
	SkillCmd.Flags().StringVarP(&scope, "scope", "s", "", "Scope of the skill")
	SkillCmd.Flags().StringVarP(&category, "category", "c", "", "Category of the skill")
	SkillCmd.Flags().StringVarP(&purpose, "purpose", "p", "", "Purpose of the skill")
	SkillCmd.Flags().IntVarP(&level, "level", "l", 0, "Level of the skill")
	SkillCmd.Flags().IntVarP(&priority, "priority", "P", 0, "Priority of the skill")
	SkillCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "Tags of the skill")
	SkillCmd.Flags().StringVarP(&file, "from-file", "f", "", "Manifest file with skill details")

	// SkillCmd.MarkFlagRequired("name")
	// SkillCmd.MarkFlagRequired("description")
	// SkillCmd.MarkFlagRequired("scope")
	// SkillCmd.MarkFlagRequired("category")
	// SkillCmd.MarkFlagRequired("purpose")
	// SkillCmd.MarkFlagRequired("level")
	// SkillCmd.MarkFlagRequired("priority")
}
