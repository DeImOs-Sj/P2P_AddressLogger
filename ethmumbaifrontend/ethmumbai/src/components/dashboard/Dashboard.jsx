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
                  Recent Peers Listed on Address Book.
                </CardDescription>
              </div>
              <Button asChild size="sm" className="ml-auto gap-1">
                <Link href="#">
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
                    <TableHead className="hidden xl:table-column">
                      Type
                    </TableHead>
                    <TableHead className="hidden xl:table-column">
                      Status
                    </TableHead>
                    <TableHead className="hidden xl:table-column">
                      Date
                    </TableHead>
                    <TableHead className="text-right">PeerId</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  <TableRow>
                    <TableCell>
                      <div className="font-medium">f5c6ba6bc3f516767c2aca7a0726b5199a2f656ae78a2b95acecd828ce2</div>
                      {/* <div className="hidden text-sm text-muted-foreground md:inline">
                        liam@example.com
                      </div> */}
                    </TableCell>
                    <TableCell className="hidden xl:table-column">
                      Sale
                    </TableCell>
                    <TableCell className="hidden xl:table-column">
                      <Badge className="text-xs" variant="outline">
                        Approved
                      </Badge>
                    </TableCell>
                    <TableCell className="hidden md:table-cell lg:hidden xl:table-column">
                      2023-06-23
                    </TableCell>
                    <TableCell className="text-right">1</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell>
                      <div className="font-medium">a1c6ba363gdfc3f516763v2aca7a0726b5199a2f656ae78a2b95acecd82</div>
                      {/* <div className="hidden text-sm text-muted-foreground md:inline">
                        olivia@example.com
                      </div> */}
                    </TableCell>
                    <TableCell className="hidden xl:table-column">
                      Refund
                    </TableCell>
                    <TableCell className="hidden xl:table-column">
                      <Badge className="text-xs" variant="outline">
                        Declined
                      </Badge>
                    </TableCell>
                    <TableCell className="hidden md:table-cell lg:hidden xl:table-column">
                      2023-06-24
                    </TableCell>
                    <TableCell className="text-right">2</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell>
                      <div className="font-medium">wo4iruj0wejf0982u37a0726b519fsdsdf9a2f656asdfdsfe78a2b95ace</div>
                      {/* <div className="hidden text-sm text-muted-foreground md:inline">
                        noah@example.com
                      </div> */}
                    </TableCell>
                    <TableCell className="hidden xl:table-column">
                      Subscription
                    </TableCell>
                    <TableCell className="hidden xl:table-column">
                      <Badge className="text-xs" variant="outline">
                        Approved
                      </Badge>
                    </TableCell>
                    <TableCell className="hidden md:table-cell lg:hidden xl:table-column">
                      2023-06-25
                    </TableCell>
                    <TableCell className="text-right">3</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell>
                      <div className="font-medium">oi20fcmow37a0726b519fsdsdf9a2f656asdsfe78fsfagowspfmm2b95</div>
                      {/* <div className="hidden text-sm text-muted-foreground md:inline">
                        emma@example.com
                      </div> */}
                    </TableCell>
                    <TableCell className="hidden xl:table-column">
                      Sale
                    </TableCell>
                    <TableCell className="hidden xl:table-column">
                      <Badge className="text-xs" variant="outline">
                        Approved
                      </Badge>
                    </TableCell>
                    <TableCell className="hidden md:table-cell lg:hidden xl:table-column">
                      2023-06-26
                    </TableCell>
                    <TableCell className="text-right">4</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell>
                      <div className="font-medium">eiposejp3qr0mp6b519fsdsdf9a2f656asdsfe78fsfagowspfmm2b3r</div>
                      {/* <div className="hidden text-sm text-muted-foreground md:inline">
                        liam@example.com
                      </div> */}
                    </TableCell>
                    <TableCell className="hidden xl:table-column">
                      Sale
                    </TableCell>
                    <TableCell className="hidden xl:table-column">
                      <Badge className="text-xs" variant="outline">
                        Approved
                      </Badge>
                    </TableCell>
                    <TableCell className="hidden md:table-cell lg:hidden xl:table-column">
                      2023-06-27
                    </TableCell>
                    <TableCell className="text-right">5</TableCell>
                  </TableRow>
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
