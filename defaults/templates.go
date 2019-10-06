package defaults

var Templates = map[string][]byte{
	"/templates/layouts/html.js": []byte(`import React from 'react';
import ReactDOM from 'react-dom';
import Head from './head';

const title = "Home | Plenti";
const heading = "Welcome to plenti!";
const desc = "Your HTML page has been hydrated and you've enabled React with Webpack and Babel!";

ReactDOM.hydrate(
  <html>
   <Head title={title} />
    <body>
      <h1>{heading}</h1>
      <p>{desc}</p>
      <p><a href="/about">About us</a>.</p>
    </body>
  </html>,
	document.getElementsByTagName('html')[0]
);`),
	"/templates/layouts/head.js": []byte(`import React, { Component } from 'react';

class Head extends Component {
  render() {
    return <head>
      <meta charset="utf-8" />
      <meta name="viewport" content="width=device-width" />
      <title>{this.props.title}</title>
    </head>
  }
}

export default Head;`),
	"/templates/layouts/header.js": []byte(`<Thing>
	Placeholder..
</Thing>`),
	"/templates/layouts/footer.js": []byte(`<Thing>
	Placeholder..
</Thing>`),
	"/templates/routed/pages.js": []byte(`<Thing>
	Placeholder..
</Thing>`),
}