package lexer

import (
	"fmt"
	"strings"

	"github.com/philipjohanszon/go-db/datastructures"
)

type Token struct {
	Command  string
	Keywords []string
}

var (
	Commands = [5]string{"GRAB", "REMOVE", "CHANGE", "MAKE", "WHERE"}
)

func Scan(data *string) (*[]Token, error) {
	splitQuery := strings.Split(*data, " ")

	queryList := convertArrayToList(&splitQuery)

	if queryList.Next == nil {
		return nil, fmt.Errorf("No query found")
	}

	//Get the first token immediately for optimisation, otherwise id have to check if it was the first index every time
	//it loops
	tokens, extractionError := extractTokens(&queryList)

	if extractionError != nil {
		fmt.Println("Error extracting tokens @lexer.Scan: ", extractionError.Error())
		return tokens, extractionError
	}

	return tokens, nil
}

func convertArrayToList(data *[]string) datastructures.StringList {
	var list datastructures.StringList = datastructures.StringList{
		Next: nil,
		End:  nil,
	}

	for _, value := range *data {
		list.AddNode(value)
	}

	return list
}

func extractTokens(dataList *datastructures.StringList) (*[]Token, error) {
	var tokens []Token = []Token{}

	//as long as the length of data is greater than 1 tokens will be extracted
	for (*dataList).Next != nil {
		token, getTokenError := getToken(dataList)

		if getTokenError != nil {
			fmt.Println("Error getting tokens @lexer.extractTokens: ", getTokenError.Error())
			return &tokens, getTokenError
		}

		tokens = append(tokens, *token)
	}

	return &tokens, nil
}

func getToken(dataList *datastructures.StringList) (*Token, error) {
	var token Token = Token{}

	// the first must be a command
	// ugly pointer dereference check
	if commandInString(&(*dataList).Next.Value) {
		//the first command is found and now we must find the "keywords" that come after, for example GRAB user (user is the "keyword" and GRAB is the command)
		token.Command = (*dataList).Next.Value
		(*dataList).RemoveOneNode()

		//to find the keywords we just loop through and stop grabbing once there is a command
		for (*dataList).Next != nil {
			//if there is a new command in the string value of the node then it will return the token because a new command is found
			if commandInString(&(*dataList).Next.Value) {
				return &token, nil
			} else {
				//if there isnt a new command the string will be added as a "keyword" and the node will be removed and then it will continue the loop
				token.Keywords = append(token.Keywords, (*dataList).Next.Value)
				(*dataList).RemoveOneNode()
			}

		}

	} else {
		return &token, fmt.Errorf("Keyword not found in the first position of the token @lexer.getToken, string: " + (*dataList).Next.Value + " " + (*dataList).Next.Next.Value)
	}

	return &token, nil
}

func commandInString(testString *string) bool {
	for _, keyword := range Commands {
		if *testString == keyword {
			return true
		}
	}
	return false
}
