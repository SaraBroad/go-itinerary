import React, { useState } from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Itinerary from "../types";

interface ListProps {
  itineraries: Itinerary[];
}

const itineraries: Itinerary[] = [
  {
    id: "12345",
    name: "Summer 2022",
    items: ["ice cream", "mini golf", "fishing"],
    startDate: "07/01/2022",
    endDate: "07/11/2022",
    locations: ["Stone Harbor, NJ"],
  },
  {
    id: "54321",
    name: "Summer 2023",
    items: ["Berlin Wall ", "Charles Bridge ", "Currywurst"],
    startDate: "08/01/2021",
    endDate: "08/15/2021",
    locations: ['Berlin', 'Prague'],
  },
  {
    id: "932923",
    name: "Summer 2024",
    items: ["TBD"],
    startDate: "06/15/2024",
    endDate: "06/27/2024",
    locations: [],
  },
];

export const ItineraryList = () => {
  const [itineraryList, setItineraryList] = useState(itineraries);

  // export const ItineraryList = ({ itinerary }: ListProps ) => {
  return (
    <div>
        <TableContainer component={Paper}>


        </TableContainer>
      {itineraryList.map((i) => {
        return (
          <div>
            <div>
              <h3>Name: {i.name}</h3>{" "}
            </div>
            <div>{i.startDate}</div>
            <div>{i.endDate}</div>
            <div>Activities: {i.items}</div>
            {i.locations?.length === 0 ? (
              <div>Locations: No destinations yet</div>
            ) : (
              <div>Locations:{i.locations}</div>
            )}
          </div>
        );
      })}
    </div>
  );
};
