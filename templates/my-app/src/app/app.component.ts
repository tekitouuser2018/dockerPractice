import { Component } from '@angular/core';
// import { MatTableModule } from '@angular/material';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
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

}
