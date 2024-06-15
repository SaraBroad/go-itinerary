import React, { useEffect, useState } from "react";
import Itinerary from "../../types";
import { ItineraryList } from "../../components/ItineraryList/ItineraryList";
import { itineraryData } from "../../data";
import axios from "axios";

export const ItineraryListPage = () => {
  // export const ItineraryListPage = ({ itineraries }: {itineraries: Itinerary[] }) => {
  // const [itineraryList, setItineraryList] = useState({})  
  const [itineraryList, setItineraryList] = useState([])  

  type GetUsersResponse = {
    data: Itinerary[];
  };

  useEffect(() => {
    itinerary()
    // getUsers()
    // GET request using fetch inside useEffect React hook
    // fetch('http://localhost:8080/itinerary/84fadc21-c07a-470c-8a7e-4113b6193a29')
    //     .then(response => response.json())
    //     .then(response => console.log(response))
        // .then(data => setItineraryList(data));

// empty dependency array means this effect will only run once (like componentDidMount in classes)
}, []);
  
  // async function getUsers() {
  //   try {
  //     // 👇️ const data: GetUsersResponse
  //     const { data, status } = await axios.get<GetUsersResponse>(
  //       'http://localhost:8080/itinerary/84fadc21-c07a-470c-8a7e-4113b6193a29',
  //       {
  //         headers: {
  //           Accept: 'application/json',
  //         },
  //       },
  //     );
  //     console.log("here")
  //     console.log("DATA", data)
  //     setItineraryList(response.json())
  //     // console.log(JSON.stringify(data, null, 4));
  
  //     // 👇️ "response status is: 200"
  //     console.log('response status is: ', status);
  
  //     return data;
  //   } catch (error) {
  //     if (axios.isAxiosError(error)) {
  //       console.log('error message: ', error.message);
  //       return error.message;
  //     } else {
  //       console.log('unexpected error: ', error);
  //       return 'An unexpected error occurred';
  //     }
  //   }
  // }
  const itinerary = async () => {
    const response = await fetch('http://localhost:8080/itineraries')

    // const response = await fetch('http://localhost:8080/itinerary/84fadc21-c07a-470c-8a7e-4113b6193a29')
    setItineraryList(await response.json())
    console.log(response.json)
  }
  

  // getUsers()

  console.log(JSON.stringify(itineraryList));
  return (
    <>
      <h3 className="itinerary-header">Itineraries</h3>
      {<ItineraryList itineraries={itineraryList} />}
      {/* <ItineraryList itineraries={itineraryData} /> */}
      {/* <ItineraryList /> */}
    </>
  );
};
