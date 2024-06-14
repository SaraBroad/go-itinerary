import { useState } from "react";
import React from "react";
import axios, { AxiosResponse } from "axios";
import { ItineraryForm } from "../../components/ItineraryForm/ItineraryForm";
import Itinerary from "../../types";
import { json } from "stream/consumers";

export const ItineraryFormPage = () => {
  const [values, setValues] = useState({
    itinerary_name: "",
    startDate: "",
    endDate: "",
    locations: [],
    // add isCompleted
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setValues({
      ...values,
      [e.target.name]: e.target.value,
    });
  };

  const options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Access-Control-Allow-Headers": "Content-Type",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "POST, PUT, PATCH, GET, DELETE, OPTIONS",
    },
    body: JSON.stringify(values),
  };

  // function saveFormData() {
  //   axios
  //     .post("http://localhost:8080/itinerary", {
  //       values,

  //       // withCredentials: true,
  //       // Headers: {
  //       //   "Content-Type": "application/json",
  //       //   "Access-Control-Allow-Headers": "Content-Type",
  //       //   "Access-Control-Allow-Origin": "*",
  //       //   "Access-Control-Allow-Methods":
  //       //     "POST, PUT, PATCH, GET, DELETE, OPTIONS",
  //       //   "Access-Control-Allow-Credentials": true,
  //       // },
  //     })
  //     .catch((err) => {
  //       console.log(err);
  //     });
  //   // const response: AxiosResponse = await axios.post(
  //   //   "http://localhost:8080/itinerary",
  //   //   dataToSend,
  //   //   {
  //   //     headers: {
  //   //       "Content-Type": "application/json",
  //   //       Accept: "application/json"
  //   //     }
  //   //   }
  //   // );

  //   // const responseData: Itinerary = response.data;
  // }

  async function saveFormData () {

  }
  const handleSubmit = (e: { preventDefault: () => void }) => {
    alert(JSON.stringify(values, null, 2));

    // console.log("values", values);
    // const dataToSend = values;

    // const response: AxiosResponse = await axios.post(
    //   "http://localhost:8080/itinerary",
    //   dataToSend
    // );
    // console.log("here");

    // const responseData: Itinerary = response.data;
    // console.log("responseData", responseData);
    // return dataToSend;
    // try {
    //   console.log('here')
    //   await saveFormData()
    // } catch (error) {
    //   console.log("error", error);
    // }

   fetch("http://localhost:8080/itinerary", {
      method: 'POST', 
      mode: 'no-cors',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(values)
    }).then(async response => setValues(await response.json()))    
    // saveFormData();
  };

  return (
    <>
      <div>
        <form className="form" onSubmit={handleSubmit}>
          <div className="formGroup">
            <label className="formLabel">Itinerary Name:</label>
            <input
              className="formField"
              type="text"
              name="itinerary_name"
              required={true}
              value={values.itinerary_name}
              onChange={handleChange}
            />
          </div>
          <div className="formGroup">
            <label className="formLabel">Start Date:</label>
            <input
              className="formField"
              type="text"
              name="startDate"
              value={values.startDate}
              onChange={handleChange}
            />
          </div>
          <div className="formGroup">
            <label className="formLabel">End Date:</label>
            <input
              className="formField"
              type="text"
              name="endDate"
              value={values.endDate}
              onChange={handleChange}
            />
          </div>
          <div className="formGroup">
            <label className="formLabel">Locations:</label>
            <input
              className="formField"
              type="text"
              name="locations"
              value={values.locations}
              onChange={handleChange}
            />
          </div>
          <button className="formSubmitBtn" type="submit">
            Submit
          </button>
        </form>
      </div>
      <ItineraryForm itinerary={values} />
    </>
  );
};
