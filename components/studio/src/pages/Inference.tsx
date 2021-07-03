import BackToHomepage from "../components/BacktoHompage"
import ThreeColumnLayout from '../layouts/ThreeColumns'
import StepsIndicator from '../components/Inference/Steps'
import RequestForm from '../components/Inference/RequestForm'
import AboutRepository from "../components/Inference/About"

export default function Inference() {
    return (
      <div className="min-h-screen bg-gray-100">
        <BackToHomepage />
        <main className="py-10">
        <ThreeColumnLayout left={<StepsIndicator/>} middle={<RequestForm/>} right={<AboutRepository/>} />
        </main>
      </div>

    )
  }