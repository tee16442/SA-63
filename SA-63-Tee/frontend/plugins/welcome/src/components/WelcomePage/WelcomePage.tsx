import React, { useState, useEffect } from 'react';
import { Alert } from '@material-ui/lab';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import {
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  makeStyles,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import { EntProblemtype } from '../../api';
import { EntCustomer } from '../../api';
import { EntRoom } from '../../api';
// header css
const HeaderCustom = {
  minHeight: '50px',
};
// css style
const useStyles = makeStyles(theme => ({
  margin: {
    margin: theme.spacing(3),
    
  },
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    margin: theme.spacing(3),
    width: 500,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));
 
export default function Create() {
 const api=new DefaultApi();
 const classes = useStyles();
 const [Problemtype, setProblemtype] = React.useState<EntProblemtype[]>([]);
 const [Customer, setCustomer] = React.useState<EntCustomer[]>([]);
 const [Room, setRoom] = React.useState<EntRoom[]>([]);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);


 
  useEffect(() => {
    const getProblemtype = async () => {
      const res = await api.listProblemtype({ limit: 3, offset: 0 });
      setProblemtype(res);
    };
  
   
    const getCustomer = async () => {
       const res = await api.listCustomer({ limit: 2, offset: 0 });
       setCustomer(res);
     };
  
     const getRoom = async () => {
      const res = await api.listRoom({ limit: 2, offset: 0 });
      setRoom(res);
    };
    getProblemtype();
    getCustomer();
    getRoom();
  }, []);

  const problemdetailhandleChange = (event:any) => {
    setproblemdetail(event.target.value as string);
  };
  const problemtypehandleChange = (event:any) => {
    setproblemtype(event.target.value as string);
  };
  const customerhandleChange = (event:any) => {
    setcustomer(event.target.value as string);
  };
  const roomhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setroom(event.target.value as number);
  };


 const [problemdetail1, setproblemdetail] = useState(String);
 const [problemtype1, setproblemtype] = useState(String);
 const [customer1, setcustomer] = useState(String);
 const [room1, setroom] = useState(Number);

let problemdetail=problemdetail1
let problemtype=Number(problemtype1)
let name=Number(customer1)
let roomnumber=Number(room1)

 const createProblem = async () => {
  let problem = {
    problemdetail,
    problemtype,
    name,
    roomnumber,
  };

  console.log(problem);
  const res: any = await api.createProblem({ problem });

  setStatus(true);
   if (res.id != ''){
     setAlert(true);
   } else {
     setAlert(false);
   }


};


 return (
  <Page theme={pageTheme.home}>
  <Header
    style={HeaderCustom} title={`ระบบแจ้งปัญหาการใช้ห้อง`}>
  </Header>
  <Content>
    <ContentHeader title="ใบเเจ้งปัญหา">
   
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 แจ้งปัญหาเรียบร้อย
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 แจ้งปัญหาไม่สำเร็จ
               </Alert>
             )}
           </div>
         ) : null}
       

    </ContentHeader>
       <div className={classes.root}>
         <form noValidate autoComplete="off">

         <div style = {{ textAlign : "center"}}>
           <FormControl className={classes.formControl}>
            <InputLabel id="demo-simple-select-label">ผู้เเจ้ง</InputLabel>
            <Select
              
              labelId="demo-simple-select-label"
              id="demo-simple-select"
              value={name || ''}
              onChange={customerhandleChange}
              
            >
              {Customer.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.nAME}
                      </MenuItem>
                    );
                  })}
            </Select>
           </FormControl>
          </div>

           <div style = {{ textAlign : "center"}}>
           <FormControl className={classes.formControl}>
            <InputLabel id="demo-simple-select-label">ห้อง</InputLabel>
            <Select
              
              labelId="demo-simple-select-label"
              id="demo-simple-select"
              value={roomnumber || ''}
              onChange={roomhandleChange}
              
            >
               {Room.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.rOOMNUMBER}
                      </MenuItem>
                    );
                  })}
              
            </Select>
           </FormControl>
          </div>

          <div style = {{ textAlign : "center"}}>
           <FormControl className={classes.formControl}>
            <InputLabel id="demo-simple-select-label">ประเภทของปัญหา</InputLabel>
            <Select
              
              labelId="demo-simple-select-label"
              id="demo-simple-select"
              value={problemtype || ''}
              onChange={problemtypehandleChange}
              
            >
               {Problemtype.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.pROBLEMTYPE}
                      </MenuItem>
                    );
                  })}
              
            </Select>
           </FormControl>
          </div>

          <br></br>
          <div style = {{ textAlign : "center"}}>

          <FormControl
                className={classes.formControl}
                variant="outlined"
          >
          <TextField
          id="outlined-multiline-static"
          label="รายละเอียดของปัญหา"
          value={problemdetail || ''}
          onChange={problemdetailhandleChange}
          multiline
          rows={6}
          defaultValue=""
          variant="outlined"
        />
        </FormControl>
        </div>

        <div style = {{ textAlign : "center"}}>
           <div className={classes.margin}>
             <Button
               onClick={() => {
                 createProblem();
               }}
               variant="contained"
               color="primary"
             >
               แจ้งปัญหา
             </Button>
           </div>
           </div>
         </form>
       </div>
     </Content>
   </Page>
 );
}