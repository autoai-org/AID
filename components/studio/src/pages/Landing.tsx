import Connection from './Connection'
import { useAppSelector, useAppDispatch } from '../services/store/hooks'
import SolversColumn from '../components/Solvers'
import Account from '../components/Account'
import Activity from '../components/Activity'

export default function Landing() {
  const server = useAppSelector((state) => state.connectivity.server)
  const dispatch = useAppDispatch()
  if (server!=="") {
    return (
      <div className="flex-grow w-full max-w-7xl mx-auto xl:px-8 lg:flex">
                    {/* Left sidebar & main wrapper */}
                    <div className="flex-1 min-w-0 bg-white xl:flex">
                        <Account />
                        <SolversColumn></SolversColumn>
                    </div>
                    <Activity />
                </div>
    )
  } else {
    return (
      <Connection handler={dispatch}></Connection>
    )
  }
  
}