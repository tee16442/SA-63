import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import MenuIcon from '@material-ui/icons/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2'; // alert
import {
    AppBar,
    Toolbar,
    Typography,
    IconButton,
    Grid,
    Container,
    FormControl,
    InputLabel,
    Select,
    TextField,
    MenuItem,
    Button
} from '@material-ui/core';

// css style 
const useStyles = makeStyles(theme => ({
    root: {
        flexGrow: 1,
    },
    textField: {
        marginLeft: theme.spacing(1),
        marginRight: theme.spacing(1),
        width: '25ch',
        color: "blue"
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    title: {
        flexGrow: 1,
    },
    container: {
        display: 'flex',
        flexWrap: 'wrap',
    },
    paper: {
        marginTop: theme.spacing(2),
        marginBottom: theme.spacing(2),
    },
    formControl: {
        width: 455,
    },
    selectEmpty: {
        marginTop: theme.spacing(2),
    },
    buttonSty: {
        margin: 'auto',
        display: 'block',
        maxWidth: '100%',
        maxHeight: '100%',
        marginBottom: 50
    }

}));

import { DefaultApi } from '../../api/apis';


import { EntCustomer } from '../../api/models/EntCustomer'; //import interface Customer 
import { EntRoom } from '../../api/models/EntRoom'; //import interface Room
import { EntAdult } from '../../api/models/EntAdult'; //import interface Adult
import { EntKid } from '../../api/models/EntKid'; //import interface Kid  
import { EntRoomamount } from '../../api/models/EntRoomamount'; //import interface Roomamount

interface Books {
    customer: number;
    room: number;
    adult: number;
    kid: number;
    roomamount: number;
    checkin: String;
    checkout: String;

}


const Books: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();

    const [books, setBooks] = React.useState<Partial<Books>>({});
    const [customers, setCustomers] = React.useState<EntCustomer[]>([]);
    const [rooms, setRooms] = React.useState<EntRoom[]>([]);
    const [adults, setAdults] = React.useState<EntAdult[]>([]);
    const [kids, setKids] = React.useState<EntKid[]>([]);
    const [roomamounts, setRoomamounts] = React.useState<EntRoomamount[]>([]);

    // alert setting
    const Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: toast => {
            toast.addEventListener('mouseenter', Swal.stopTimer);
            toast.addEventListener('mouseleave', Swal.resumeTimer);
        },
    });

    const getCustomer = async () => {
        const res = await api.listCustomer({ limit: 10, offset: 0 });
        setCustomers(res);
    };

    const getRoom = async () => {
        const res = await api.listRoom({ limit: 10, offset: 0 });
        setRooms(res);
    };

    const getAdult = async () => {
        const res = await api.listAdult({ limit: 10, offset: 0 });
        setAdults(res);
    };

    const getKid = async () => {
        const res = await api.listKid({ limit: 10, offset: 0 });
        setKids(res);
    };
    const getRoomamount = async () => {
        const res = await api.listRoomamount({ limit: 10, offset: 0 });
        setRoomamounts(res);
    };

    // Lifecycle Hooks
    useEffect(() => {
        getRoom();
        getAdult();
        getKid();
        getRoomamount();
        getCustomer();

    }, []);

    // set data to object books
    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: unknown }>,
    ) => {
        const name = event.target.name as keyof typeof books;
        const { value } = event.target;
        setBooks({ ...books, [name]: value });
        console.log(books);
    };

    // clear input form
    function clear() {
        setBooks({});
    }

    // function save data
    function save() {
        books.checkin +=":00+07:00";
        books.checkout +=":00+07:00";
        const apiUrl = 'http://localhost:8080/api/v1/bookss';
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(books),
        };

        console.log(books); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

        fetch(apiUrl, requestOptions)
            .then(response => {
                console.log(response)
                response.json()
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
            })
    }
   
    return (

        
        <div className={classes.root}>
            <AppBar position="static">
                <Toolbar>
                    <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
                        <MenuIcon />
                    </IconButton>
                    <Typography variant="h6" className={classes.title}>
                        ระบบจองห้อง
          </Typography>
                    <div>
                        <IconButton
                            aria-label="account of current user"
                            aria-controls="menu-appbar"
                            aria-haspopup="true"
                            color="inherit"
                        >
                            <AccountCircle />
                        </IconButton>
                    </div>
                </Toolbar>
            </AppBar>
            <Container maxWidth="sm">

                <Grid container spacing={3}>

                    <Grid item xs={10}>
                        <h2 > จองห้องพัก </h2>
                    </Grid>


                    <Grid item xs={12}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel >เลือกประเภทห้องพัก</InputLabel>
                            <Select
                                name="room"
                                value={books.room || ''}
                                onChange={handleChange}
                                label="เลือกประเภทห้องพัก"
                            >
                                {rooms.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.roomtype}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel >เลือกจำนวนผู้เข้าพัก(ผู้ใหญ่)</InputLabel>
                            <Select
                                name="adult"
                                value={books.adult || ''}
                                onChange={handleChange}
                                label="เลือกจำนวนผู้เข้าพัก(ผู้ใหญ่)"
                                fullWidth
                            >
                                {adults.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.amount}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel>เลือกจำนวนผู้เข้าพัก(เด็ก)</InputLabel>
                            <Select
                                name="kid"
                                value={books.kid || ''}
                                onChange={handleChange}
                                label="เลือกจำนวนผู้เข้าพัก(เด็ก)"
                            >
                                {kids.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.amount}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel>เลือกจำนวนห้องพัก</InputLabel>
                            <Select
                                name="roomamount"
                                value={books.roomamount || ''}
                                onChange={handleChange}
                                label="เลือกจำนวนห้องพัก"
                            >
                                {roomamounts.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.amount}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>

                    <Grid item xs={12}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel >ยืนยันผู้จอง</InputLabel>
                            <Select
                                name="customer"
                                value={books.customer || ''}
                                onChange={handleChange}
                                label="ยืนยันผู้จอง"
                                fullWidth
                            >
                                {customers.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.email}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>


                    <Grid item xs={12}>
                        <form className={classes.container} noValidate>
                            <TextField
                                label="เลือกวันเข้าพัก"
                                id="checkin"
                                name="checkin"
                                type="datetime-local"
                                value={books.checkin || ''} // (undefined || '') = ''
                                className={classes.textField}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                onChange={handleChange}
                            />
                        </form>
                    </Grid>

                    <Grid item xs={12}>
                        <form className={classes.container} noValidate>
                            <TextField
                                label="เลือกวันออกจากห้องพัก"
                                id="checkout"
                                name="checkout"
                                type="datetime-local"
                                value={books.checkout || ''} // (undefined || '') = ''
                                className={classes.textField}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                onChange={handleChange}
                            />
                        </form>
                    </Grid>

                    <Grid item xs={10}>
                        <Button
                            name="saveData"
                            size="large"
                            variant="contained"
                            color="primary"
                            disableElevation
                            className={classes.buttonSty}
                            onClick={save}
                        >
                        บันทึกการจอง
              </Button>
                    </Grid>
                </Grid>
            </Container>
        </div>
    );
};

export default Books;