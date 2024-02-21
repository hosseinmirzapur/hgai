# Using Google Gemini Model within Go Application

This application uses Google's Gemini LLM to respond to any user prompt.

## Features

- This application is deployed on `Vercel Platform` via usage of `Vercel Serverless Functions`.

- The Go version of `Langchain` is used for the development.

- Purely written in Go with no external libraries used rather than Langchain for its huge support over LLMs.

- An easy to use interface


## Contributions

Feel free to open PRs and issues on this repository, any suggestions are appreciated.


## License

MIT

## Endpoints available

- Base URL: https://hgai-three.vercel.app/api

- [POST] /generative/prompt

````js

# Request Body

{
    "prompt": "some text in here", // required
}

````