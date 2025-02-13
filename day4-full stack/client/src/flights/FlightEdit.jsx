import { useNavigate } from "react-router-dom";
import PageHeader from "../header/PageHeader";

function FlightEdit() {
    const navigate = useNavigate()
    const onGo = function() {
        navigate('/flights/list')
    }
    // change happening through code
    return (
        <>
            <PageHeader  PageNumber={1}/>
            <h3><a href="/flights/list" className="btn btn-light">Go Back</a>Edit Flight Ticket Price</h3> {/*change happening through hyperlink*/}
            <button class = "btn btn-success" 
                onClick={onGo}>Go back using code</button> {/*change happening through code*/} 
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="number" className="form-label">Flight Number:</label>
                    <div className="form-control" id="number">???</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="airline_name" className="form-label">Airline Name:</label>
                    <div type="text" className="form-control" id="airline_name">???</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="source" className="form-label">Source:</label>
                    <div className="form-control" id="source">???</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="destination" className="form-label">Destination:</label>
                    <div className="form-control" id="destination">???</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="capacity" className="form-label">Capacity(Number of Seats):</label>
                    <div className="form-control" id="capacity">???</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="price" className="form-label">Ticket Price:</label>
                    <input type="text" className="form-control" id="price" placeholder="Please enter ticket price" />
                </div>
                <button className="btn btn-warning">Update Price</button>
            </div>
        </>
    );
}

export default FlightEdit;