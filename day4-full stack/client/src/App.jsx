import FlightList from "./flights/FlightList"
import FlightCreate from "./flights/FlightCreate"
import FlightEdit from "./flights/FlightEdit"

import { BrowserRouter, Route, Routes } from 'react-router-dom'
import FullName from "./FullName"
import PageHeader from "./header/PageHeader"
function App() {
  return (
    <>     
      <BrowserRouter>
        <Routes>
          <Route path="/flights/list" element={<FlightList/>}/>
          <Route path="/flights/create" element={<FlightCreate/>}/>
          <Route path="/flights/edit/:id" element={<FlightEdit/>}/>
        </Routes>
      </BrowserRouter>
      
    </>
  )
}

export default App
