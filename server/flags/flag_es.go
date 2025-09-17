package flags

import (
	"bufio"
	"fmt"
	"os"
	"server/model/elasticsearch"
	"server/service"
)

func Elasticsearch() error {
	esService := service.ServiceGroupApp.EsService

	indexExists, err := esService.IndexExists(elasticsearch.ArticleIndex())

	if err != nil {
		return err
	}

	if indexExists {
		fmt.Println("The index already exists. Do you want to delete current data and recreate the index?(y/n)")

		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "y":
			// 如果用户输入 y，删除索引
			fmt.Println("Proceeding to delete the data and recreate the index...")
			if err := esService.IndexDelete(elasticsearch.ArticleIndex()); err != nil {
				return err
			}
		case "n":
			// 如果用户输入 n，退出程序
			fmt.Println("Exiting the program.")
			os.Exit(0)
		default:
			// 如果用户输入无效，提示重新输入
			fmt.Println("Invalid input. Please enter 'y' to delete and recreate the index, or 'n' to exit.")
			return Elasticsearch() // 递归调用，重新输入
		}
	}
	return esService.IndexCreate(elasticsearch.ArticleIndex(), elasticsearch.ArticleMapping())
}
