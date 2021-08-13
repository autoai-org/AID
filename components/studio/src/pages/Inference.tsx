import React from 'react';

import BackToHomepage from "../components/BacktoHompage"
import ThreeColumnLayout from '../layouts/ThreeColumns'
import StepsIndicator from '../components/Inference/Steps'
import RequestForm from '../components/Inference/RequestForm'
import AboutRepository from "../components/Inference/About"
import Result from '../components/Inference/Result'
import Analysis from '../components/Inference/Analysis'
import { restclient } from '../services/apis'
import { ALL_SOLVERS } from '../services/apis/queries'
interface InferenceState {
  result: object;
  finished: boolean;
  analysis: boolean;
  currentStep: string;
  solver: object
}

export default class Inference extends React.Component<any, InferenceState> {
  constructor(props: any) {
    super(props)
    this.state = {
      result: {},
      finished: false,
      analysis: false,
      currentStep: 'Form Request',
      solver: {}
    }
  }
  componentDidMount() {
    let containerID = this.props.match.params.containerID
    let self = this
    restclient.query(ALL_SOLVERS).then(function(res:any) {
      res.data.data.solvers.map(function(each:any){
        if (each.image.container.uid === containerID) {
          self.setState({solver: each})
        }
      }) 
    })
  }
  handleResponse = (res: any) => {
    this.setState({ result: res })
    this.setState({ finished: true })
    this.setState({ currentStep: 'View Results' })
  }
  handleAnalysisPage = () => {
    this.setState({ analysis: true })
    this.setState({ currentStep: 'Detailed Analysis' })
  }
  render() {
    return (
      <div className="min-h-screen bg-gray-100">
        <BackToHomepage />
        <main className="py-10">
          {!this.state.finished &&
            <ThreeColumnLayout left={<StepsIndicator currentStep={this.state.currentStep} />} middle={<RequestForm onReceivedResponse={this.handleResponse} />} right={<AboutRepository solver={this.state.solver}/>} />
          }
          {this.state.finished && !this.state.analysis &&
            <ThreeColumnLayout left={<StepsIndicator currentStep={this.state.currentStep} />} middle={<Result onNextStep={this.handleAnalysisPage} src={this.state.result} />} right={<AboutRepository solver={this.state.solver} />} />
          }
          {this.state.analysis && this.state.analysis &&
            <ThreeColumnLayout left={<StepsIndicator currentStep={this.state.currentStep} />} middle={<Analysis src={this.state.result} />} right={<AboutRepository solver={this.state.solver} />} />
          }
        </main>
      </div>

    )
  }
}