import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { Layout } from 'antd';
import Home from './components/home/home';

const { Content } = Layout;
const WHITE = '#ffffff';

class App extends Component {


  render() {
    return (
      <Layout style={{ padding: '24px' }}>
        <Content 
          style={{
            margin: 0,
            minHeight: 'calc(100vh - 48px)',
            backgroundColor: WHITE
          }}
        >
          <Router>
              <Switch>
                <Route path='/' component = {(Home)} />
              </Switch>
          </Router>
        </Content>
      </Layout>
    )
  }
}

export default App


// export default App;
