import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { UsersService } from './users.service';


const routes: Routes = [];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
  providers:[UsersService]
})
export class AppRoutingModule { }
