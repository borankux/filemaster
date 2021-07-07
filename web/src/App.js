import './App.css';
import {BrowserRouter as Router, Link, Route} from "react-router-dom";
import React, {Component} from "react";


class  App extends Component{
    state = {
        gists: null
    }

    componentDidMount() {
        fetch("https://api.github.com/gists")
            .then(res => res.json())
            .then(gists => {
                this.setState({gists})
            })
    }

    render() {
        const {gists} = this.state
        return (
            <div className="App">
                <Router>
                <div className="info-box">
                    <div className="info">
                        <h3>Scanners</h3>
                        <hr/>
                        <span>1</span>
                    </div>
                    <ul>
                        {gists ? (gists.map(gist => (
                            <li key={gist.id}>
                                <Link to={`/g/${gist.id}`}>
                                    {gist.description}
                                </Link>
                            </li>
                            ))) :
                            (<div>loading...</div>)}
                    </ul>
                </div>

                <div className="page">
                    {gists && (
                        <Route path="/g/:gistId" render={({match}) => (
                        <Gist gist={gists.find(g => g.id === match.params.gistId)} />
                        )}/>
                    )}
                </div>
                </Router>
            </div>
        );
    }
}


const Gist = ({ gist }) => {
    return (
        <div>
            <h3>{gist.description}</h3>
            <div>
            </div>
        </div>
    )
}

export default App;
