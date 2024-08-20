# gorelease

Simple fullstack application using sveltekit in the FE and Golang in the BE to maintain
visual history of Golang releases.

## Description

The application runs on Deta Space cloud platform.
The FE send the BE request with code to run and BE Golang server then pass the request to
[Golang Playground](https://go.dev) and gets the response for the code and returns it to the FE.
FE then displays the output of the program.

## Running locally

Make sure that you have already installed [The space CLI](https://deta.space/docs/en/build/fundamentals/space-cli/) from [Deta Space](https://deta.space/)

```bash
# Run locally
space dev
```

## Topics covered

- To show how to use sveltekit with Golang Backend server.
- Allow to run the code directly from the browser.
- Deployment of fullstack application on Deta Space.
