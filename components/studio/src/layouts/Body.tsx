import {
    BrowserRouter as Router,
    Switch,
    Route,
} from "react-router-dom";

import AIDHeader from '../components/Header'
import Landing from '../pages/Landing'
import Detail from '../pages/Detail'
import Inference from '../pages/Inference'

export default function MainBody() {
    return (
        <>
            {/* Background color split screen for large screens */}
            <div className="fixed top-0 left-0 w-1/2 h-full bg-white" aria-hidden="true" />
            <div className="fixed top-0 right-0 w-1/2 h-full bg-gray-50" aria-hidden="true" />
            <div className="relative min-h-screen flex flex-col">

                <AIDHeader />
                <Router>
                    <Switch>
                        <Route path="/details/:solverID" component={Detail}/>
                        <Route path="/inference">
                            <Inference />
                        </Route>
                        <Route path="/">
                            <Landing />
                        </Route>
                    </Switch>
                </Router>
            </div>
        </>
    )
}