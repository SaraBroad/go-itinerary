import React, { ReactHTMLElement, useCallback, useState } from "react";
import Button from "@mui/material/Button";

export const ItineraryForm = () => {
  const [name, setName] = useState("");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");
  const [locations, setLocations] = useState<string[]>([]);

  const [formData, setFormData] = useState({
    name: "",
    startDate: "",
    endDate: "",
    locations: [],
  });

  const handleLocationChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value;
    setLocations([...locations, value]);
  };

  const handleSubmit = (e: { preventDefault: () => void }) => {
    alert("Form Submitted");
    console.log(e);
  };

  return (
    <>
    <div className="form-container">
      <form onSubmit={handleSubmit} className="form-fields"> 
        <label>
          Name{": "}
          <input value={name} onChange={(e) => setName(e.target.value)} />
        </label>
        <label>
          StartDate{": "}
          <input
            value={startDate}
            onChange={(e) => setStartDate(e.target.value)}
          />
        </label>
        <label>
          End Date{": "}
          <input value={endDate} onChange={(e) => setEndDate(e.target.value)} />
        </label>
        <label>
          Locations{": "}
          <input value={locations} onChange={handleLocationChange} />
        </label>
        <input type="submit" value="Submit" />
      </form>

      {/* <Button variant="contained" onClick={handleSubmit}>
        Submit
      </Button> */}
        </div>
    </>
  );
};
