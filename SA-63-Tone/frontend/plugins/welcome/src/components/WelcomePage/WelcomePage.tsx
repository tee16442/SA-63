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
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis/'; // Api Gennerate From Command
import {
  EntCustomer,
  EntEmployee,
  EntPaymenttype,
  EntRoomtype,
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

interface payment {
  customer: Number,
  employee: Number,
  payday: String,
  roomType: Number,
  paymenttype: Number,
}

const Payment: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  
  const [payment, setPayment] = React.useState<
    Partial<payment>
  >({});

  const [customer, setCustomer] = React.useState<EntCustomer[]>([]);
  const [roomType, setRoomtype] = React.useState<EntRoomtype[]>([]);
  const [paymenttype, setPaymenttype] = React.useState<EntPaymenttype[]>([]);
  const [employee, setEmployee] = React.useState<EntEmployee[]>([]);

 
  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setEmployee(res);
  };

  const getCustomer = async () => {
    const res = await http.listCustomer({ limit: 10, offset: 0 });
    setCustomer(res);
  };

  const getRoomtype = async () => {
    const res = await http.listRoomtype({ limit: 10, offset: 0 });
    setRoomtype(res);
  };

  const getPaymenttype = async () => {
    const res = await http.listPaymenttype({ limit: 10, offset: 0 });
    setPaymenttype(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getEmployee();
    getCustomer();
    getPaymenttype();
    getRoomtype();
  }, []);

  // set data to object payment
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Payment;
    const { value } = event.target;
    setPayment({ ...payment, [name]: value });
    console.log(payment);
  };

  // clear input form
  function clear() {
    setPayment({});
  }

  // function save data
  function save() {
    payment.payday += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/payments';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payment),
    };

    // alert setting
    const Toast = Swal.mixin({
      toast: true,
      position: 'top-end',
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
    });
    
    console.log(payment); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => {
        console.log(response.json());
        if (response.ok === true) {
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
      <Header style={HeaderCustom} title={`ระบบชำระเงิน`}>
      </Header>

      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อ User ลูกค้า</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือก User Name</InputLabel>
                <Select
                  name="customer"
                  value={payment.customer || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {customer.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.uSERNAME}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ยอดเงินที่ต้องชำระ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกราคาห้องพัก</InputLabel>
                <Select
                  name="roomType"
                  value={payment.roomType || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {roomType.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.rOOMPRICE}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ช่องทางการชำระ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกประเภทการชำระ</InputLabel>
                <Select
                  name="paymenttype"
                  value={payment.paymenttype || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {paymenttype.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.tYPE}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อพนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกพนักงาน</InputLabel>
                <Select
                  name="employee"
                  value={payment.employee || ''} // (undefined || '') = ''
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

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่ชำระเงิน</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกวันที่"
                  name="payday"
                  type="datetime-local"
                  value={payment.payday || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
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
                บันทึกการชำระเงิน
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default Payment;
