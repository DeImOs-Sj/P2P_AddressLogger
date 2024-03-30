import React, { useState, useEffect } from 'react';
import axios from 'axios';

const LatestBlock = () => {
  const [blockData, setBlockData] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('http://127.0.0.1:7000/v1/latest_block');
        setBlockData(response.data);
      } catch (error) {
        setError(error);
      }
    };

    fetchData();
  }, []);

  if (error) {
    return <div>Error: {error.message}</div>;
  } else if (!blockData) {
    return <div>Loading...</div>;
  } else {
    return (
      <div>
        <h1>Latest Block Data</h1>
        <pre>{JSON.stringify(blockData, null, 2)}</pre>
      </div>
    );
  }
};

export default LatestBlock;
