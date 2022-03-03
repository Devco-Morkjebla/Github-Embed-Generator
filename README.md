#### Go API project for generating svg for github embed

## Routes
Languages are comma seperated
> /skills/:languages
Example: `https://arvidgithubembed.herokuapp.com/skills?languages=php,mysql,javascript,typescript`
### Customization

- `titlecolor` - Card's title color _(hex color)_
- `textcolor` - Body text color _(hex color)_
- `bordercolor` - Card's border color _(hex color)_.
- `backgroundcolor` - Card's background color _(hex color)_
- `title` - Card's custom title _(string)_
- `boxcolor` - Color of the boxes with the logos inside _(hex color)_



 users are comma seperated
> /rankList/:users

#### Common Options:
All hex colors without '#' please
- `titlecolor` - Card's title color _(hex color)_
- `textcolor` - Body text color _(hex color)_
- `bordercolor` - Card's border color _(hex color)_.
- `backgroundcolor` - Card's background color _(hex color)_ 
- `boxcolor` - Card's languages color _(hex color)_
- `title` - Card's custom title _(string)_
Example: `https://arvidgithubembed.herokuapp.com/skills?languages=php,mysql,javascript,typescript&title=test`




Example: 
`/ranklist?users=lartrax,arvidwedtstein&bordercolor=black&titlecolor=red&textcolor=green&backgroundcolor=yellow&title=test`


> `/mostactivity?org=devco-morkjebla`
// https://api.github.com/repos/arvidwedtstein/nuxt-website/stats/code_frequency
//https://mholt.github.io/json-to-go/


//https://awesome-github-stats.azurewebsites.net/user-stats/lartrax/rank
//https://api.github.com/repos/devco-morkjebla/unzippy/stats/commit_activity