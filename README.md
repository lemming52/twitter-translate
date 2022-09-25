# Twitter Translation

![](docs/media/babel.png)

## Team Members
Andrew: https://github.com/lemming52

## Tool Description
This tool is a fairly classic example of an i-was-lazy-and-thought-a-tool-would-be-faster tool. I'm unsure exactly if the time save of this is worth the effort of this code.

The main idea for this tool came from when I was doing some geolocation work for the Russian advances in the 2022 Russo-Ukrainian war. I found that in trying to find tweets about a particular town or location, I'd get different results based on if I'd search for latin or cyrillic versions of place names i.e Nedryhailiv/Недригайлів or Kiev/Kyiv/Київ.

This tool was basically me being lazy; I wanted to search for the same thing in multiple languages, without managing the back and forth from google translate to twitter and back and forward. This seemed small enough in scope for me to work on myself.

The tool is designed to consist of two components, a backend API that manages the translation requests and conversion to links, and a Chrome Extension (this seemed simplest to me, other options are available) that would allow a user to type in a search term, select the languages and execute the searches.

In terms of points from the spreadsheet, I viewed this as useful as it seemed others had wanted thematically similar tools:

> A tool that helps find the right search terms

> A tool that suggests different ways of spelling a name

> Online databases for terms in different languages

### *HACKATHON NOTES*

This tool is for all intents and purposes, unusable for anyone but me in this form. With that in mind, please don't bother spending much time on grading this, all that would be interesting or useful is if this seems useful for anyone but me. I advise against trying to install it, and suggest just reading this doc and the [brief mocked guide](https://github.com/lemming52/twitter-translate/blob/master/demo.md).

## Installation
At the time of writing, installation for this tool is very difficult and not guaranteed to work. The hope was an end user would only have to install the chrome extension.

A significant extra complication for the installation is that the translation is currently built off of the Google Cloud Platform Translation API, which is not something we can reasonably expect open-source investigators to have access to or to configure.

Given the whole thing doesn't function to completion, I wouldn't bother trying to install it.

If you definitely want to, you'll need [Golang 1.17](https://go.dev/doc/install). Then:
```
go mod download
make build
./twitter-translate
```

In addition, you'll need chrome for the extension part, but that's not working meaningfully atm so there's little point.

## Usage

The envisioned usage was as follows:
1. User clicks chrome extension icon
2. User enters search term, selects languages from permitted subset, hits execute
3. User has tabs launched per language, can analyse as required.

To show what we have currently, please look at the [demonstration doc](https://github.com/lemming52/twitter-translate/blob/master/demo.md), but be aware that this is intended to demonstrate the functionality in lieu of installing, and also to highlight the current barely existent nature of the tool. If interested, there's also a short video that can be provided on request.

## Additional Information

### Design Limitations

This tool is definitely beyond the complexity we can reasonably expect an OSINT Researcher to make use of, without completely separating out and making publicly accessible the backend.

That presents a pretty obvious limitation as most translation APIs have cost limitations for public usage and also the hosting costs of the service have to be borne, for my testing / hackathon use case, I was happy to back this off my own.

My memory of dealing with them is poor since it was years ago I last looked into them, but publicly accessible Translation APIs are all IIRC limited in terms of usage rate and price, so trying to build a tool that makes use of them will always run into problems.

If we could manage all of those however, the installation cost for a researcher is minimal, in a chrome extension.

### Extensions

There's a lot of potential for extensions, either small scale additions to my baseline functionality here or bigger features that are thematically similar.

#### Alternatives

Here the basic mode is to add translations for a given text string. The sheet mentioned alternative variants of names as an option, so rather than just language translation, modes could be added to the API to:

* a `Names` mode, search for alternatives to a person name
* a general term mode, search for say `SU-25/MIG-29/SU-34/SU-35`

If this pattern of quick searches in other languages is useful, it's readily applicable to other websites outside of twitter.

#### User-refined content

While we're basing this on google translate ATM, if required user-created specific alternatives could be added to the stored translations, i.e Kiev and Kyiv, which isn't quite a translation.

#### Additional tooling.

Little bit beyond but conceptually an OSINT chrome extension that allows for tooling direct from twitter or other sites, i.e. integration with the auto-capture spreadsheet.

#### Translation options

This currently hard codes without flexibility Google Translate. It could be extended to use other translation engines, and to use user accounts to authenticate to the Translate API instead of mine.

#### Partial translation

It's possible that a user doesn't want to translate all the query, just sections. This is an easy addition.

### Further development

In a word; lots.

* The backend needs to have a mature deployment setup so that it's accessible not just locally and uses a specific provisioned GCP account or other hosting system. Getting this set up proved to be to much to ask in the time I provisioned for this work.
* The extension currently is non functional, the popup/extension javascript needs writing to be easily useable and to work; a web form with text boxes and language search/dropdown for example. Frontend Javascript is not something I have more than brief exposure to.
* Probably needs to be a bit more security consciousness than I've added here.
* Get the docker-composition working, as currently it doesn't load the google translator.