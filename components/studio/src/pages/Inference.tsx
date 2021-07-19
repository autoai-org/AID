import React from 'react';

import BackToHomepage from "../components/BacktoHompage"
import ThreeColumnLayout from '../layouts/ThreeColumns'
import StepsIndicator from '../components/Inference/Steps'
import RequestForm from '../components/Inference/RequestForm'
import AboutRepository from "../components/Inference/About"
import Result from '../components/Inference/Result'
import Analysis from '../components/Inference/Analysis'

interface InferenceState {
  result: object;
  finished: boolean;
  analysis: boolean;
  currentStep: string;
}

export default class Inference extends React.Component<any, InferenceState> {
  constructor(props: any) {
    super(props)
    this.state = {
      result: {},
      finished: false,
      analysis: false,
      currentStep: 'Form Request',
    }
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
            <ThreeColumnLayout left={<StepsIndicator currentStep={this.state.currentStep} />} middle={<RequestForm onReceivedResponse={this.handleResponse} />} right={<AboutRepository />} />
          }
          {this.state.finished && !this.state.analysis &&
            <ThreeColumnLayout left={<StepsIndicator currentStep={this.state.currentStep} />} middle={<Result onNextStep={this.handleAnalysisPage} src={this.state.result} />} right={<AboutRepository />} />
          }
          {this.state.analysis && this.state.analysis &&
            <ThreeColumnLayout left={<StepsIndicator currentStep={this.state.currentStep} />} middle={<Analysis src={this.state.result} />} right={<AboutRepository />} />
          }
        </main>
      </div>

    )
  }
}