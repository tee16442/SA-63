import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
Content,
Header,
Page,
pageTheme,
ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select'; 

import { EntRoomtype } from '../../api';
import { EntRoomstate } from '../../api';
import { EntPromotion } from '../../api';

const api=new DefaultApi();
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
        marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
      width: '25ch',
    },
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
 }),
);

 export default function Create() {
  const profile = { givenName  : 'G17' };
  const classes = useStyles();

  const [Roomtype, setRoomtype] = React.useState<EntRoomtype[]>([]);
  const getRoomtype = async () => {
    const res = await api.listRoomtype({ limit: 3, offset: 0 });
    setRoomtype(res);
  };
  useEffect(() => {
    getRoomtype();
  }, []);
  const [Roomstate, setRoomstate] = React.useState<EntRoomstate[]>([]);
  const getRoomstate = async () => {
    const res = await api.listRoomstate({ limit: 2, offset: 0 });
    setRoomstate(res);
  };
  useEffect(() => {
    getRoomstate();
  }, []);
  const [Promotion, setPromotion] = React.useState<EntPromotion[]>([]);
  const getPromotion = async () => {
    const res = await api.listPromotion({ limit: 3, offset: 0 });
    setPromotion(res);
  };
  useEffect(() => {
    getPromotion();
  }, []);



  const promotiondetailhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setpromotiondetail(event.target.value as number);
  };
  const roomnumberhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setroomnumber(event.target.value as string);
  };
  const roomstatehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setroomstate(event.target.value as number);
  };
  const typedetailhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    settypedetail(event.target.value as number);
  };


  const [roomnumber, setroomnumber] = useState(String);
  const [promotiondetail1, setpromotiondetail] = useState(Number);
  const [roomstate1, setroomstate] = useState(Number);
  const [typedetail1, settypedetail] = useState(Number);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  let promotion = Number(promotiondetail1)
  let roomstate = Number(roomstate1)
  let roomtype = Number(typedetail1)

  const room = {
    promotion,
    roomnumber,
    roomstate,
    roomtype,
  };
  console.log(room);
  const createRoom = async () => {
    const res: any = await api.createRoom({ room });
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
      title={`ระบบเพิ่มข้อมูลห้องพัก ${profile.givenName || 'to Backstage'}`}
      subtitle="รายละเอียดข้อมูลของห้องพัก"
    ></Header>
    <Content>
      <ContentHeader title="เพิ่มข้อมูลห้องพัก">
        {status ? (
          <div>
            {alert ? (
              <Alert severity="success">
              เพิ่มข้อมูลห้องเรียบร้อย
              </Alert>
            ) : (
              <Alert severity="warning" style={{ marginTop: 20 }}>
              เพิ่มข้อมูลห้องไม่สำเร็จ
              </Alert>
            )}
          </div>
        ) : null}
      </ContentHeader>
      <div className={classes.root}>
        <form noValidate autoComplete="off">

        <FormControl 
            fullWidth
            variant="outlined"
            className={classes.formControl}
        >
            <InputLabel></InputLabel>
            <div
            
            
            >
            <TextField 
                  label="หมายเลขห้อง" 
                  name = "rOOMNUMBER"
                  type="string"
                  defaultValue=" " 
                  value={roomnumber || ''}
                  onChange={roomnumberhandleChange}
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  />
          </div>
        </FormControl>

        <FormControl
            fullWidth
            variant="outlined"
            className={classes.formControl}
        >
            <InputLabel>ประเภทห้อง</InputLabel>
            <Select
            name="tYPEDEATAIL"
            value={roomtype || ''}
            onChange={typedetailhandleChange}>
            {Roomtype.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.tYPEDEATAIL}
                      </MenuItem>
                    );
                  })}
          </Select>
        </FormControl>

        <FormControl 
            fullWidth
            variant="outlined"
            className={classes.formControl}
        >
            <InputLabel>สถานะห้อง</InputLabel>
            <Select
            name="ROOMSTATE"
            value={room.roomstate || ''}
            onChange={roomstatehandleChange}>
            {Roomstate.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.rOOMSTATE}
                      </MenuItem>
                    );
                  })}
          </Select>
        </FormControl>

        <FormControl 
            fullWidth
            variant="outlined"
            className={classes.formControl}
        >
            <InputLabel>โปรโมชั่น</InputLabel>
            <Select
            name="pROMOTIONDETAIL"
            value={promotion || ''}  
            onChange={promotiondetailhandleChange}>
            {Promotion.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.pROMOTIONDETAIL}
                      </MenuItem>
                    );
                  })}
          </Select>
        </FormControl>

          <div className={classes.margin}>
            <Button
            onClick={() => {
              createRoom();
            }}
            variant="contained"
            color="primary"
            >
            เพิ่ม
          </Button>
          <Button
            style={{ marginLeft: 20 }}
            component={RouterLink}
            to="/"
            variant="contained"
          >
            กลับ
          </Button>
        </div>
      </form>
    </div>
  </Content>
</Page>
);
}
