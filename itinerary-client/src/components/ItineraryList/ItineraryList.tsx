import React, { useState } from "react";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import Itinerary from "../../types";
import "../itinerary.css"

interface ListProps {
  itineraries: Itinerary[];
}

export const ItineraryList = ({itineraries}: ListProps) => {
console.log("itineraries", itineraries)
  const [itineraryList, setItineraryList] = useState(itineraries);
  return (
    <div>
      {/* <TableContainer component={Paper}>Hello</TableContainer> */}
      <div className="list-component">
      {itineraryList.map((i) => {
        return (
          <div>
            <div>
              <h3>Name: {i.name}</h3>{" "}
            </div>
            <div>Start Date:{i.startDate}</div>
            <div>End Date:{i.endDate}</div>
            {i.items?.length === 0 ? (
              <div>Activities: Nothing to do yet</div>
            ) : (
              <div>Activities:{i.items}</div>
            )}
            {/* <div>Activities: {i.items}</div> */}
            {i.locations?.length === 0 ? (
              <div>Locations: No destinations yet</div>
            ) : (
              <div>Locations:{i.locations}</div>
            )}
          </div>
        );
      })}
      </div>
    </div>
  );
};
