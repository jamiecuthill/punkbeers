# Punk Beers

An exercise for Golang Training using the Punk (Brewdog) API.

[https://punkapi.com/documentation/v2](https://www.google.com/url?q=https://punkapi.com/documentation/v2&sa=D&ust=1568186931290000&usg=AOvVaw2MWv4e_o1MdQxffa2Bptbl)

We'll build this as a CLI script using [https://github.com/spf13/cobra](Cobra).

1. The primary goal of your application is to return the strongest beer by abv you can find in the API. For example I would issue the command in my terminal

        $ punkbeers strongest

2. Return the strongest n beers where n is a user supplied parameter of your application.

        $ punkbeers strongest -n 5

3. Return the strongest beer for a given food pairing, either do the matching yourself or use the food param of the beers api.

        $ punkbeers strongest -food pork
