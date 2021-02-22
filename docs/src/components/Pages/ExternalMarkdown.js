import React, { Component } from 'react'
import ReactMarkdown from 'react-markdown'

class Markdown extends Component {
  constructor(props) {
    super(props)
    this.props = props
    this.state = { mdfile: null }
  }

  componentDidMount() {
    fetch(this.props.src).then((response) => response.text()).then((text) => {
      this.setState({ mdfile: text })
    })
  }

  render() {
    return (
      <div className="content">
        <ReactMarkdown source={this.state.mdfile} />
      </div>
    )
  }
}

export default Markdown