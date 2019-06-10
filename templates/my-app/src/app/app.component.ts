import { Component, Input } from '@angular/core';
import { UsersService } from './users.service';
import { User } from '../model';
import { Observable } from 'rxjs';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(private usersService: UsersService){}

  // @Input() usersService: UsersService
  // response$ : Observable<User[]>;
  response : User[];
  // ngOnInit(){
  //   this.response$ = this.usersService.getUsers();
  // }

  responseHeader =[
    'id',
    'name',
    'created_time',
    'updated_time',
  ];

  title = 'my-app';
  contentsHeader =[
    'id',
    'name',
  ];
  contents =[
    {id:'1',name:'firstBrits'},
    {id:'2',name:'econdMan'},
    {id:'3',name:'3rd'},
  ];

  onSubmit (){
    // this.response =[
    //   {id:1,name:'firstBrits',created_time:'2019/06/10',updated_time:'2019/06/10'},
    //   {id:2,name:'econdMan',created_time:'2019/06/10',updated_time:'2019/06/10'},
    //   {id:3,name:'3rd',created_time:'2019/06/10',updated_time:'2019/06/10'},
    // ];
    this.usersService.getUsers()
    .subscribe(
      (res) => { 
        this.response = res["TableResponse"];
        console.log(res);
        // console.log(res["TableResponse"]);
        // console.log(this.response[0]);
        // console.log(this.response[1]);
        alert('fetched data');
      },
      error =>{
        alert('feiled fetching data');
        console.log('error: ', error);
      }
      
    );


   }

}
