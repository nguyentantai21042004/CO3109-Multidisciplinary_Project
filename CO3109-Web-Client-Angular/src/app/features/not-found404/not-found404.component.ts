import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-not-found404',
  templateUrl: './not-found404.component.html',
  styleUrls: ['./not-found404.component.css']
})
export class NotFound404Component {
  constructor(private router: Router) { }

  goHome() {
    this.router.navigate(['/']);
  }
}
