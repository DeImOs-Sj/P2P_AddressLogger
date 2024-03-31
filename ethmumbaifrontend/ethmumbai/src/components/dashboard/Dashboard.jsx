import { Link } from 'react-router-dom';
import React, { useState, useEffect } from 'react';
import axios from 'axios';

import {
  Activity,
  ArrowUpRight,
  CircleUser,
  CreditCard,
  DollarSign,
  Menu,
  Package2,
  Search,
  Users,
} from "lucide-react"

import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/avatar"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Input } from "@/components/ui/input"
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"

export function Dashboard() {
  const [blockData, setBlockData] = useState(null);
    const [versionData, setVersionData] = useState({});
  const [statusData, setStatusData] = useState({});

  const [mode, setMode] = useState('');
  const [loading, setLoading] = useState(true);
    const [addressBookData, setAddressBookData] = useState([]);

console.log(addressBookData)
              // console.log(blockData)

  const [error, setError] = useState(null);


//latest data
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
  //mode
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('http://127.0.0.1:7000/v1/mode');
        setMode(response.data);
        setLoading(false);
      } catch (error) {
        setError('Error fetching mode data');
        setLoading(false);
      }
    };

    fetchData();
  }, []);
  //v2 version 
  useEffect(() => {
    const fetchVersion = async () => {
      try {
        const response = await axios.get('http://127.0.0.1:7000/v2/version');
        setVersionData(response.data);
        setLoading(false);
      } catch (error) {
        setError('Error fetching version data');
        setLoading(false);
      }
    };

    fetchVersion();
  }, []);
  
  //v2 status 
  useEffect(() => {
    const fetchStatus = async () => {
      try {
        const response = await axios.get('http://127.0.0.1:7000/v2/status');
        const { modes, genesis, network } = response.data;
        setStatusData({ modes, genesis, network });
        setLoading(false);
      } catch (error) {
        setError('Error fetching status data');
        setLoading(false);
      }
    };

    fetchStatus();
  }, []);

  //address book 
  useEffect(() => {
    const fetchAddressBook = async () => {
      try {
        const response = await axios.get('http://127.0.0.1:8000/v1/address_book');
        const data = response.data;
        const jsonData = JSON.parse("[" + data.replace(/}{/g, "},{") + "]");
        const filteredData = jsonData.map(obj => {
          const peerID = Object.keys(obj)[0];
          const multiaddr = obj[peerID][0][1];
          return { peerID, multiaddr };
        });
        setAddressBookData(filteredData);
        console.log(filteredData);
        setLoading(false);
      } catch (error) {
        setError('Error fetching address book data');
        setLoading(false);
      }
    };

    fetchAddressBook();
  }, []);
  
  return (
    <div className="flex min-h-screen w-full flex-col">
    
      <main className="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
        <div className="grid gap-4 md:grid-cols-2 md:gap-8 lg:grid-cols-4">
          <Card>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">
                V1 Latest Block Number 
              </CardTitle>
              <DollarSign className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">607658</div>
              <p className="text-xs text-muted-foreground">
The latest block processed by light client.
              </p>
        
            </CardContent>
          </Card>
          <Card>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">
                V1 Mode
              </CardTitle>
              <Users className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">{ mode}</div>
 <p className="text-xs text-muted-foreground">
Retrieves the operating mode of the Avail light client. The Light client can operate in two different modes, LightClient or AppClient

              </p>
            </CardContent>
          </Card>
          <Card>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">V2 Version </CardTitle>
              <CreditCard className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>

              <div className="text-2xl font-bold">{versionData.version}</div>
              <p className="text-xs text-muted-foreground">
                Gets the version of the light client binary, and the version of the compatible network.


              </p>
            </CardContent>
          </Card>
          <Card>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-sm font-medium">V2 Status</CardTitle>
              <Activity className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">Modes: {statusData.modes}</div>
              <p className="text-xs text-muted-foreground ">Genesis Hash:0x6f09966420b2608d1947ccfb0f
                2a362450d1fc7fd902c29b67c906eaa965a7ae</p>
              
              <p className='text-xs text-muted-foreground '>Network: {statusData.network}</p>
            </CardContent>
          </Card>
        </div>
        <div className="grid gap-4 md:gap-8 lg:grid-cols-2 xl:grid-cols-3">
          <Card className="xl:col-span-2">
            <CardHeader className="flex flex-row items-center">
              <div className="grid gap-2">
                <CardTitle>Address Book</CardTitle>
        <CardDescription>
  {/* Recent Peers Listed on Address Book.
{posts.length && posts.map((post, index) => (
    <li key={index}>
      <p>Name: {entry.name}</p>
      <p>Email: {entry.email}</p>
    </li>
  ))} */}
</CardDescription>

              </div>
      <Button asChild size="sm" className="ml-auto gap-1">
                <Link to="#">
                  View All
                  <ArrowUpRight className="h-4 w-4" />
                </Link>
              </Button>
            </CardHeader>
            <CardContent>
             <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead>Peer Public key </TableHead>
                    <TableHead className="text-right">Multiaddr</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {addressBookData.map((entry, index) => (
                    <TableRow key={index}>
                      <TableCell>
                        <div className="font-medium">{entry.peerID}</div>
                      </TableCell>
                      <TableCell className="text-right">{entry.multiaddr}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </CardContent>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Latest Block Information</CardTitle>
            </CardHeader>
            <CardContent className="grid gap-8">
              <div className="flex items-center gap-4">
              
                <div className="grid gap-1">
                  <p className="text-sm font-medium leading-none">
                    Block Number
                  </p>
                  <p className="text-sm text-muted-foreground">
607658                  </p>
                </div>
              </div>
              <div className="flex items-center gap-4">
               
                <div className="grid gap-1">
                  <p className="text-sm font-medium leading-none">
                    Hash
                  </p>
                  <p className="text-sm text-muted-foreground">
                    0xa91fe3282b15875e265ebe297c7d94184e6301c125f1449
                    ad50699c822789b9c                  </p>
                </div>
              </div>
              <div className="flex items-center gap-4">
                
                <div className="grid gap-1">
                  <p className="text-sm font-medium leading-none">
                    Parent_Hash
                  </p>
                  <p className="text-sm text-muted-foreground">
                    0x34456e054d7c5f943736e35ee285ab4ef37bc30705fa5c
                    e5cddc42bd68a3f667                  </p>
                </div>
              </div>
              <div className="flex items-center gap-4">
               
                <div className="grid gap-1">
                  <p className="text-sm font-medium leading-none">
                    State_Root
                  </p>
                  <p className="text-sm text-muted-foreground">
                    0x804187e7b3805b421249c4aed14bf423b9b737328824f7c8c9a9ceff1ba63093
                  </p>
                </div>
              </div>
              <div className="flex items-center gap-4">
                
                <div className="grid gap-1">
                  <p className="text-sm font-medium leading-none">
                    Extrinsics_Root
                  </p>
                  <p className="text-sm text-muted-foreground">
                    0x0e56c8d4514b91d42ab14c6ed4090a423d2dc4741e104309b7757dfe1a407774
                  </p>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </main>
    </div>
  )
}
