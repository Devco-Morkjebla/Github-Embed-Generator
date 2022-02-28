#### Go API project for generating svg for github embed

## Routes
> /skills/:languages
Example: `https://arvidgithubembed.herokuapp.com/skills?languages=php,mysql,javascript,typescript`
### Customization

You can customize the appearance of your `Card` however you wish with URL params.

#### Common Options:

- `titlecolor` - Card's title color _(hex color)_
- `textcolor` - Body text color _(hex color)_
- `bordercolor` - Card's border color _(hex color)_.
- `backgroundcolor` - Card's background color _(hex color)_
- `boxcolor` - Card's languages color _(hex color)_
- `title` - Card's custom title _(string)_
Example: `https://arvidgithubembed.herokuapp.com/skills?languages=php,mysql,javascript,typescript&title=test`

# users are comma seperated
> /rankList/:users


### Customization

You can customize the appearance of your `Card` however you wish with URL params.

#### Common Options:

- `titlecolor` - Card's title color _(hex color)_
- `textcolor` - Body text color _(hex color)_
- `bordercolor` - Card's border color _(hex color)_.
- `backgroundcolor` - Card's background color _(hex color)_
- `title` - Card's custom title _(string)_


Example: 
`/ranklist?users=lartrax,arvidwedtstein&bordercolor=black&titlecolor=red&textcolor=green&backgroundcolor=yellow&title=test`


> `/mostactivity?org=devco-morkjebla`
/repository 
// https://api.github.com/repos/arvidwedtstein/nuxt-website/stats/code_frequency
//https://mholt.github.io/json-to-go/

//http://localhost:8080/mostactivity?org=devco-morkjebla&bordercolor=ff0000&boxcolor=222222&backgroundcolor=000000&titlecolor=ffffff&textcolor=dddddd&textcolor=000000
//https://arvidgithubembed.herokuapp.com/ranklist?users=lartrax,arvidwedtstein,alvaage,migliusmockus