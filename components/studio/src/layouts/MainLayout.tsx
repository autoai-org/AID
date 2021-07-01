import Landing from '../pages/Landing'
import ReactDOM from "react-dom";
import {
    BrowserRouter as Router,
    Switch,
    Route
  } from "react-router-dom";
export default function MainLayout() {
    return (
        <div className="mx-auto">
            <main><Landing/></main>
        </div>
    )
}