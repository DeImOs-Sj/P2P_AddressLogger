import { Dashboard } from "./components/dashboard/Dashboard"
// import Providers from "./components/Providers/Providers"
import {Routes,Route} from "react-router-dom"
import Chart from "./components/CreateChart/Chart"
import Navbar from "./components/Navbar/Navbar"
import LatestBlock from "./components/LatestBlock/Lastest"
function App() {
   // Sample peer data
  const chartData = [
    { peerId: 1, peerPublicKey: 'publicKey1', timestamp: '2022-01-30T10:00:00' },
    { peerId: 2, peerPublicKey: 'publicKey2', timestamp: '2021-03-30T11:00:00' },
    { peerId: 1, peerPublicKey: 'publicKey3', timestamp: '2024-03-30T12:00:00' },
    { peerId: 3.11, peerPublicKey: 'publicKey3', timestamp: '2024-03-30T12:00:00' },
    { peerId: 4, peerPublicKey: 'publicKey3', timestamp: '2024-03-30T12:00:00' },
    { peerId: 2.8, peerPublicKey: 'publicKey3', timestamp: '2024-03-30T12:00:00' },

  ];
 const options = {
    maintainAspectRatio: false,
    responsive: true,
    height: 40,
    width: 50
 };
  
  return (
    <>
      <Navbar/>
     <Routes>

        <Route path="/" element={<Dashboard />} />
                <Route
          path="/dashboard"
          element={<Dashboard />}
        />
     <Route
          path="/chart"
          element={<Chart  chart={chartData}  options={options}/>}
        />
         {/*   <Route
          path="/create_region"
          element={<Create />}
        /> */}
                {/* <Route path="/subregions" element={<SubRegions />} /> */}

       {/*  <Route path="/subregions" element={<SubRegions />} /> */}
      </Routes>


     </>
    // <>
    //   <LatestBlock/>
    // </>
  )
}

export default App
