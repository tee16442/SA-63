import React, { useState,FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button'
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';

import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { EntCustomertype } from '../../api/models/EntCustomertype';
import { EntTitle } from '../../api/models/EntTitle';
import { EntGender } from '../../api/models/EntGender';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    '& > *': {
      margin: theme.spacing(1),
    },
  },
  large: {
    width: theme.spacing(15),
    height: theme.spacing(15),
  },
  paper: {
    padding: theme.spacing(2),
    
    height: theme.spacing(53),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  },
  formControl: {
    margin: theme.spacing(1),
    minWidth: 130,
  },
  textField: {
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
    width: 200,
  },
  }));

  const WelcomePage: FC<{}> = () => {
    const classes = useStyles();
    const api=new DefaultApi();


  const [Customertype, setCustomertype] = React.useState<EntCustomertype[]>([]);
  const [Gender, setGender] = React.useState<EntGender[]>([]);
  const [Title, setTitle] = React.useState<EntTitle[]>([]);
  
  useEffect(() => {
  const getCustomertype = async () => {
    const res = await api.listCustomertype({ limit: 10, offset: 0 });
    setCustomertype(res);
  };
  const getGender = async () => {
    const res = await api.listGender({ limit: 10, offset: 0 });
    setGender(res);
  };
  const getTitle = async () => {
    const res = await api.listTitle({ limit: 10, offset: 0 });
    setTitle(res);
  };
  
  getTitle();
  getGender();
  getCustomertype();
  }, []);


  const agehandleChange = (event: React.ChangeEvent<{ value: unknown }>)=> {
    setage(event.target.value as number);
  };
  const emailhandleChange = (event: any) => {
    setemail(event.target.value as string);
  };
  const namehandleChange = (event: any) => {
    setname(event.target.value as string);
  };
  const passhandleChange = (event: any) => {
    setpass(event.target.value as string);
  };
  const usernamehandleChange = (event: any) => {
    setusername(event.target.value as string);
  };
  const genderhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setgender(event.target.value as number);
  };
  const titletypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    settitletype(event.target.value as number);
  };


  const customertypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setcustomertype(event.target.value as number);
  };

  const [age1, setage] = useState(Number);


  const [email, setemail] = useState(String);
  const [name, setname] = useState(String);
  const [password, setpass] = useState(String);
  const [username, setusername] = useState(String);
  const [gender1, setgender] = useState(Number);
  const [customertype1, setcustomertype] = useState(Number);
  const [title1, settitletype] = useState(Number);


let age = Number(age1)
let customertype = Number(customertype1)
let gender = Number(gender1)
let title = Number(title1)


  
    const customer = {
      age,
      customertype,
      email,
      gender,
      name,
      password,
      title,
      username,
    };
    console.log(customer);
    const createCustomer = async () => {
    
    const res: any = await api.createCustomer({ customer  });
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ระบบ จองห้องพัก`}
      >
      </Header>
    <Content>
        <ContentHeader title="ระบบสมาชิก">
        </ContentHeader>
    <Paper className={classes.paper}>
      <Grid container wrap="nowrap" spacing={2}>
        
            <Grid item xs zeroMinWidth>
                <TextField 
                  label="User NAME" 
                  name = "uSERNAME"
                  type="string"
                  defaultValue=" " 
                  value={username || ''}
                  onChange={usernamehandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  />
                  &nbsp;&nbsp;&nbsp;


                <TextField 
                  label="password" 
                  name = "pASSWORD"
                  type="password"
                  defaultValue=" " 
                  value={password || ''}
                  onChange={passhandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  />
                <br></br><br></br>


                <FormControl className={classes.formControl}>
                    <InputLabel>คำนำหน้า</InputLabel>
                      <Select
                        
                        name="Title"
                        value={title || ''} 
                        onChange={titletypehandleChange} 
                      > 
                    {Title.map(item => {
                        return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.tITLETYPE}
                      </MenuItem>
                         );
                       })}
                      </Select>
                </FormControl>&nbsp;&nbsp;&nbsp;


                <TextField 
                label="ชื่อ" 
                name = "nAME"
                type="string"
                defaultValue=" " 
                value={name || ''}
                onChange={namehandleChange}
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
                />
                  &nbsp;&nbsp;&nbsp;


      <FormControl className={classes.formControl}>
        <InputLabel>เพศ</InputLabel>
          <Select
            name="gender"
            value={gender || ''} 
            onChange={genderhandleChange} 
          > 
            {Gender.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.gENDER}
                      </MenuItem>
                    );
                  })}
          </Select>
      </FormControl>
      <br></br><br></br>


      <TextField 
        label="อายุ" 
        name = "age"
        type="int"
        defaultValue=" " 
        value={age || ''}
        onChange={agehandleChange}
        className={classes.textField}
        InputLabelProps={{
          shrink: true,
        }}
        />
                  &nbsp;&nbsp;&nbsp;
      <br></br><br></br>


      <TextField 
        label="E-Mail" 
        name = "eMAIL"
        type="string"
        defaultValue=" " 
        value={email || ''}
        onChange={emailhandleChange}
        className={classes.textField}
        InputLabelProps={{
          shrink: true,
        }}
        />
        &nbsp;&nbsp;&nbsp;

      </Grid>       
      </Grid>  
      <FormControl className={classes.formControl}>
        <InputLabel>ประเภทสมาชิก</InputLabel>
          <Select
            name="customertype"
            value={customertype || ''}  
            onChange={customertypehandleChange} 
          > 
              {Customertype.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.cUSTOMERTYPE}
                      </MenuItem>
                    );
                  })}
          </Select>

      </FormControl>
      <div style = {{ textAlign : "left"}}>
      <Button 
       onClick={() => {
         createCustomer();
       }}
       variant="contained"
       color="primary"
     >
          Save
      </Button> 
 
    </div> 
    </Paper>
      </Content>
    </Page>
  );

};
 
  export default WelcomePage;
