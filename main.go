package main

import (
  "context"
  "fmt"
  "log"
"flag"
  "github.com/tmc/langchaingo/llms"
  "github.com/tmc/langchaingo/llms/ollama"
)

func main() {
  var queryFlag string
  flag.StringVar(&queryFlag, "query","","prompt for search")
  flag.Parse()
  llm, err := ollama.New(ollama.WithModel("llama3"))
  if err != nil {
    log.Fatal(err)
  }

  ctx := context.Background()
  completion, err := llms.GenerateFromSinglePrompt(ctx, llm, queryFlag)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("[termprompt]: %v\n", completion)
}
