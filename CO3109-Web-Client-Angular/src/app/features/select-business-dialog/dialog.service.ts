import { Injectable } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { SelectBusinessDialogComponent } from './component/select-business-dialog.component';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class DialogService {
  constructor(private dialog: MatDialog, private router: Router) {}

  openBusinessSelectionDialog() {
    const dialogRef = this.dialog.open(SelectBusinessDialogComponent, {
      width: '600px',
      disableClose: true
    });
  
    dialogRef.componentInstance.businessSelected.subscribe(business => {
      this.router.navigate(['/dashboard', business.id]);
      dialogRef.close();
    });
  }
}