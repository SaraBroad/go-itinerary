import React, { ReactHTMLElement, useCallback, useState } from "react";
import Button from "@mui/material/Button";

export const ItineraryForm = () => {
  const [values, setValues] = useState({
    name: "",
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

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    console.log(values);
    alert(JSON.stringify(values, null, 2));
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
              name="name"
              required={true}
              value={values.name}
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
    </>
  );
};
