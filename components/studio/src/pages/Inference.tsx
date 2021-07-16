import React from 'react';

import BackToHomepage from "../components/BacktoHompage"
import ThreeColumnLayout from '../layouts/ThreeColumns'
import StepsIndicator from '../components/Inference/Steps'
import RequestForm from '../components/Inference/RequestForm'
import AboutRepository from "../components/Inference/About"
import Result from '../components/Inference/Result'

interface InferenceState {
  result: object;
  finished: boolean;
  currentStep: string;
}

export default class Inference extends React.Component<any, InferenceState> {
  constructor(props: any) {
    super(props)
    this.state = {
      result: {},
      finished: false,
      currentStep: 'Form Request',
    }
  }
  handleResponse=(res: any) => {
    this.setState({ result: res })
    this.setState({ finished: true })
    this.setState({ currentStep: 'View Results' })
  }
  render() {
    return (
      <div className="min-h-screen bg-gray-100">
        <BackToHomepage />
        <main className="py-10">
          {!this.state.finished &&
            <ThreeColumnLayout left={<StepsIndicator currentStep={this.state.currentStep}/>} middle={<RequestForm onReceivedResponse={this.handleResponse} />} right={<AboutRepository />} />
          }
          {this.state.finished &&
            <ThreeColumnLayout left={<StepsIndicator currentStep={this.state.currentStep}/>} middle={<Result src={this.state.result} />} right={<AboutRepository />} />
          }
        </main>
      </div>

    )
  }
}