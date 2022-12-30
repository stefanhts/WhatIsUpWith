import React, { useState } from 'react';
import axios from 'axios';

interface Props {
  // Props go here
}

const SearchBar: React.FC<Props> = (props) => {
  const [searchTerm, setSearchTerm] = useState('');

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(event.target.value);
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    // send a post request to localhost:8080/search
    // with the search term as the body
    // and then log the response
    //
    // hint: use axios.post

    axios.post('http://localhost:8080/search', { searchTerm , maxResults: 10})
      .then((response) => {
        // Do something with the response here
        console.log(response);
      })
      .catch((error) => {
        // Handle the error here
        console.error(error);
      });
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        Search:
        <input type="text" value={searchTerm} onChange={handleChange} />
      </label>
      <input type="submit" value="Submit" />
    </form>
  );
};

export default SearchBar;

