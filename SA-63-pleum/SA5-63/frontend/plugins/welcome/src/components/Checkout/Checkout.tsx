import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis/'; // Api Gennerate From Command
import {
  //EntBooks,
  EntCustomer,
  EntEmployee,
  EntRoom,

} from '../../api/models';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
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

interface checkout {
  employee: Number,
  customer: Number,
  books: Number,
  room: Number,
  checkoutdate: String,
}

const checkout: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [checkout, setCheckout] = React.useState< Partial<checkout>>({});

  const [customer, setCustomer] = React.useState<EntCustomer[]>([]);
  //const [books, setBooks] = React.useState<EntBooks[]>([]);
  const [employee, setEmployee] = React.useState<EntEmployee[]>([]);
  const [room, setRoom] = React.useState<EntRoom[]>([]);

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
  });

  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setEmployee(res);
  };

  const getCustomer = async () => {
    const res = await http.listCustomer({ limit: 10, offset: 0 });
    setCustomer(res);
  };

 /* const getBooks = async () => {
    const res = await http.listBooks({ limit: 10, offset: 0 });
    setBooks(res);
  };*/

  const getRoom = async () => {
    const res = await http.listRoom({ limit: 10, offset: 0 });
    setRoom(res);
  };


  // Lifecycle Hooks
  useEffect(() => {
    getEmployee();
    getCustomer();
    //getBooks();
    getRoom();
  }, []);

  // set data to object checkout
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof checkout;
    const { value } = event.target;
    setCheckout({ ...checkout, [name]: value });
    console.log(checkout);
  };

  // clear input form
  function clear() {
    setCheckout({});
  }

  // function save data
  function save() {
    checkout.checkoutdate +=":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/checkouts';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(checkout),
    };

    console.log(checkout); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบเช็คเอาท์`}>
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10 }}>Team G17</div>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อลูกค้า</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกชื่อลูกค้า</InputLabel>
                <Select
                  name="customer"
                  value={checkout.customer || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {customer.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.nAME}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>หมายเลขห้องพัก</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกหมายเลขห้อง</InputLabel>
                <Select
                  name="room"
                  value={checkout.room || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {room.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.rOOMNUMBER}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

           
            <Grid item xs={3}>
            <div className={classes.paper}>วันที่เช็คเอาท์</div>
            </Grid>
            <Grid item xs={9}>
            <form className={classes.container} noValidate>
                <TextField
                    id="datetime-local"
                    label="เลือกวันที่เช็คเอาท์"
                    type="datetime-local"
                    value={checkout.checkoutdate}
                    className={classes.textField}
                    InputLabelProps={{
                    shrink: true,
                }}
                onChange={handleChange}
            />
            </form>
            </Grid>

            <Grid item xs={3}>
            <div className={classes.paper}>ชื่อพนักงานที่บันทึกข้อมูล</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกชื่อพนักงาน</InputLabel>
                <Select
                  name="employee"
                  value={checkout.employee || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {employee.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.eMPLOYEENAME}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกการเช็คเอาท์
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default checkout;
