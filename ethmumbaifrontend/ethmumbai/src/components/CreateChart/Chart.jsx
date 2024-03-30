import React from 'react';
import { Line } from 'react-chartjs-2';
import { CategoryScale, Chart,LinearScale,PointElement,LineElement ,Filler} from "chart.js";

Chart.register(CategoryScale,LinearScale,PointElement,LineElement,Filler);

const ChartData = ({ chart }) => {
  console.log(chart)
  const peerIds = chart.map(peer => peer.peerId);
  const peerPublicKeys = chart.map(peer => peer.peerPublicKey);
  const timestamps = chart.map(peer => peer.timestamp);

  // Creating chart data
  const data = {
    labels: timestamps,
    datasets: [
      {
        label: 'Peer ID',
        data: peerIds,
        fill: false,
        borderColor: 'rgb(75, 192, 192)',
        tension: 0.1
      },
      {
        label: 'Peer Public Key',
        data: peerPublicKeys,
        fill: true,
        borderColor: 'rgb(122, 11, 10)',
        tension: 0.1,
      backgroundColor: 'rgb(122, 11, 10)'

      }
    ]
  };

  return (
    <div className='flex flex-col justify-center'> 
            <h2 className='text-3xl text-semibold ml-[48rem] m-[1rem]'>Peer Details Chart</h2>

      <Line data={data} />

    </div>
  );
};

export default ChartData;
